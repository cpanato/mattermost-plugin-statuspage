package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/pkg/errors"

	pluginapi "github.com/mattermost/mattermost-plugin-api"
	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/plugin"
)

type Plugin struct {
	plugin.MattermostPlugin

	// configuration is the active plugin configuration. Consult getConfiguration and
	// setConfiguration for usage.
	configuration *configuration

	BotUserID string
	ChannelID string

	// configurationLock synchronizes access to the configuration.
	configurationLock sync.RWMutex
}

func (p *Plugin) OnActivate() error {
	configuration := p.getConfiguration()

	if err := p.IsValid(configuration); err != nil {
		return errors.Wrap(err, "validating the configuration")
	}

	client := pluginapi.NewClient(p.API, p.Driver)
	botID, err := client.Bot.EnsureBot(&model.Bot{
		Username:    "statuspage_bot",
		DisplayName: "StatusPage",
		Description: "Created by the StatusPage plugin.",
	})
	if err != nil {
		return errors.Wrap(err, "ensuring StatusPage bot")
	}
	p.BotUserID = botID

	team, TeamErr := p.API.GetTeamByName(p.configuration.Team)
	if TeamErr != nil {
		return errors.Wrap(err, "getting team name")
	}

	channel, appErr := p.API.GetChannelByName(team.Id, p.configuration.Channel, false)
	if appErr != nil && appErr.StatusCode == http.StatusNotFound {
		channelToCreate := &model.Channel{
			Name:        p.configuration.Channel,
			DisplayName: p.configuration.Channel,
			Type:        model.ChannelTypeOpen,
			TeamId:      team.Id,
			CreatorId:   botID,
		}

		newChannel, errChannel := p.API.CreateChannel(channelToCreate)
		if errChannel != nil {
			return errors.Wrap(errChannel, "creating the channel")
		}

		p.ChannelID = newChannel.Id

		return nil
	} else if appErr != nil {
		return errors.Wrap(appErr, "getting the channel to check if that already exists")
	}

	p.ChannelID = channel.Id

	return nil
}

func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Mattermost StatusPage Plugin"))
		return
	}
	switch r.URL.Path {
	case "/webhook":
		token := r.URL.Query().Get("token")
		if token == "" || strings.Compare(token, p.configuration.Token) != 0 {
			errorMessage := "Invalid or missing token"
			p.postHTTPDebugMessage(errorMessage)
			http.Error(w, errorMessage, http.StatusBadRequest)
			return
		}

		service := r.URL.Query().Get("service")
		if service == "" {
			errorMessage := "You must provide a service name"
			p.postHTTPDebugMessage(errorMessage)
			http.Error(w, errorMessage, http.StatusBadRequest)
			return
		}

		p.handleWebhook(r.Body, service, p.ChannelID, p.BotUserID)

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	default:
		p.postHTTPDebugMessage("Invalid URL path")
		http.NotFound(w, r)
	}
}

func (p *Plugin) IsValid(configuration *configuration) error {
	if configuration.Team == "" {
		return fmt.Errorf("must set a Team")
	}

	if configuration.Channel == "" {
		return fmt.Errorf("must set a Channel")
	}

	if configuration.Token == "" {
		return fmt.Errorf("must set a Token")
	}

	return nil
}

func (p *Plugin) postHTTPDebugMessage(errorMessage string) {
	p.API.LogDebug("Failed to serve HTTP request", "Error message", errorMessage)
}
