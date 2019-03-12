package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/mattermost/mattermost-server/model"
)

const (
	statusPageIconURL  = "https://pbs.twimg.com/profile_images/963832478728314880/QoqF8Db1_400x400.jpg"
	statusPageUsername = "Status Page Notification"
)

func (p *Plugin) handleWebhook(body io.Reader, service, channelID, userID string) {
	p.API.LogInfo("Received statuspage notification", "service=", service)
	var t *StatusPageNotification
	if err := json.NewDecoder(body).Decode(&t); err != nil {
		p.postHTTPDebugMessage(err.Error())
		return
	}
	p.API.LogInfo("Message to improve statuspage", "msg=", t.ToJson())

	var color string
	var fields []*model.SlackAttachmentField
	if t.Component != nil {
		p.API.LogDebug(t.Component.Status)
		msg := fmt.Sprintf("Status: %s", t.Component.Status)
		if t.Component.Description != "" {
			msg = fmt.Sprintf("Status: %s\nDescription: %s", t.Component.Status, t.Component.Description)
		}
		fields = addFields(fields, t.Component.Name, msg, false)
		if t.ComponentUpdate != nil {
			msg := fmt.Sprintf("Old Status: %s\nNew Status: %s", t.ComponentUpdate.OldStatus, t.ComponentUpdate.NewStatus)
			fields = addFields(fields, "", msg, true)
		}
		color = setColor(t.Component.Status)
	}

	if t.Incident != nil {
		fields = addFields(fields, t.Incident.Name, t.Incident.Status, false)
		fields = addFields(fields, "Impact", t.Incident.Impact, true)
		fields = addFields(fields, "Link", t.Incident.Shortlink, true)

		createdAt := t.Incident.CreatedAt
		updatedAt := t.Incident.UpdatedAt
		fields = addFields(fields, "Created At", createdAt.String(), true)
		fields = addFields(fields, "Updated At", updatedAt.String(), true)

		for _, incidentUpdate := range t.Incident.IncidentUpdates {
			msg := fmt.Sprintf("Status: %s\nDescription: %s\nUpdatedAt: %s", incidentUpdate.Status, incidentUpdate.Body, incidentUpdate.UpdatedAt.String())
			fields = addFields(fields, "Incident Update", msg, false)
		}
		color = setColor(t.Incident.Impact)
	}

	serviceStatusName := fmt.Sprintf("%s Status - %s", strings.ToUpper(service), t.Page.StatusDescription)
	attachment := &model.SlackAttachment{
		Title:  serviceStatusName,
		Fields: fields,
		Color:  color,
	}

	post := &model.Post{
		ChannelId: channelID,
		UserId:    userID,
		Props: map[string]interface{}{
			"from_webhook":      "true",
			"override_username": statusPageUsername,
			"override_icon_url": statusPageIconURL,
		},
	}

	model.ParseSlackAttachment(post, []*model.SlackAttachment{attachment})
	if _, appErr := p.API.CreatePost(post); appErr != nil {
		p.postHTTPDebugMessage(appErr.Message)
		return
	}
}

func addFields(fields []*model.SlackAttachmentField, title, msg string, short bool) []*model.SlackAttachmentField {
	return append(fields, &model.SlackAttachmentField{
		Title: title,
		Value: msg,
		Short: short,
	})
}

func setColor(impact string) string {
	mapImpactColor := map[string]string{
		"maintenance":          "#ADD8E6",
		"operational":          "#008000",
		"degraded_performance": "#FFFF66",
		"partial_outage":       "#FFA500",
		"major_outage":         "#FF0000",
		"major":                "#FF4500",
		"minor":                "#FF6347",
	}

	if val, ok := mapImpactColor[impact]; ok {
		return val
	}

	return "#F0F8FF"
}
