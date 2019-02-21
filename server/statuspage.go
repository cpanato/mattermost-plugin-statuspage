package main

import (
	"encoding/json"
	"time"
)

type Meta struct {
	Unsubscribe   string    `json:"unsubscribe,omitempty"`
	Documentation string    `json:"documentation,omitempty"`
	GeneratedAt   time.Time `json:"generated_at,omitempty"`
}

type Page struct {
	Id                string `json:"id,omitempty"`
	StatusIndicator   string `json:"status_indicator,omitempty"`
	StatusDescription string `json:"status_description,omitempty"`
}

type Component struct {
	CreatedAt   time.Time `json:"created_at,omitempty"`
	Description string    `json:"description,omitempty"`
	ID          string    `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	PageID      string    `json:"page_id,omitempty"`
	Position    int       `json:"position,omitempty"`
	Status      string    `json:"status,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	GroupID     string    `json:"group_id,omitempty"`
	Showcase    bool      `json:"showcase,omitempty"`
}

type ComponentUpdate struct {
	OldStatus     string    `json:"old_status,omitempty"`
	NewStatus     string    `json:"new_status,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	ComponentType string    `json:"component_type,omitempty"`
	State         string    `json:"state,omitempty"`
	Id            string    `json:"id,omitempty"`
	ComponentId   string    `json:"component_id,omitempty"`
}

type IncidentUpdate struct {
	Body               string    `json:"body,omitempty"`
	CreatedAt          time.Time `json:"created_at,omitempty"`
	DisplayAt          time.Time `json:"display_at,omitempty"`
	ID                 string    `json:"id,omitempty"`
	IncidentID         string    `json:"incident_id,omitempty"`
	Status             string    `json:"status,omitempty"`
	TwitterUpdatedAt   time.Time `json:"twitter_updated_at,omitempty"`
	UpdatedAt          time.Time `json:"updated_at,omitempty"`
	WantsTwitterUpdate bool      `json:"wants_twitter-update,omitempty"`
}

type Incident struct {
	Backfilled                    bool              `json:"backfilled,omitempty"`
	Components                    *[]Component      `json:"components,omitempty"`
	CreatedAt                     time.Time         `json:"created_at,omitempty"`
	ID                            string            `json:"id,omitempty"`
	Impact                        string            `json:"impact,omitempty"`
	ImpactOverride                string            `json:"impact_override,omitempty"`
	IncidentUpdates               []*IncidentUpdate `json:"incident_updates,omitempty"`
	MonitoringAt                  time.Time         `json:"monitoring_at,omitempty"`
	Name                          string            `json:"name,omitempty"`
	PageID                        string            `json:"page_id,omitempty"`
	PostmortemBody                string            `json:"postmortem_body,omitempty"`
	PostmortemBodyLastUpdatedAt   time.Time         `json:"postmortem_body_last_updated_at,omitempty"`
	PostmortemIgnored             bool              `json:"postmortem_ignored,omitempty"`
	PostmortemNotifiedSubscribers bool              `json:"postmortem_notified_subscribers,omitempty"`
	PostmortemNotifiedTwitter     bool              `json:"postmortem_notified_twitter,omitempty"`
	PostmortemPublishedAt         time.Time         `json:"postmorem_published_at,omitempty"`
	ResolvedAt                    time.Time         `json:"resolved_at,omitempty"`
	ScheduledAutoInProgress       bool              `json:"scheduled_auto_in_progress,omitempty"`
	ScheduledAutoCompleted        bool              `json:"scheduled_auto_completed,omitempty"`
	ScheduledFor                  time.Time         `json:"scheduled_for"`
	ScheduledRemindPrior          bool              `json:"scheduled_remind_prior,omitempty"`
	ScheduledRemindedAt           time.Time         `json:"scheduled_reminded_at,omitempty"`
	ScheduledUntil                time.Time         `json:"scheduled_until,omitempty"`
	Shortlink                     string            `json:"shortlink,omitempty"`
	Status                        string            `json:"status,omitempty"`
	UpdatedAt                     time.Time         `json:"updated_at,omitempty"`
}

type IncidentResponse []Incident

type StatusPageNotification struct {
	Meta            *Meta            `json:"meta,omitempty"`
	Page            *Page            `json:"page,omitempty"`
	Component       *Component       `json:"component,omitempty"`
	ComponentUpdate *ComponentUpdate `json:"component_update,omitempty"`
	Incident        *Incident        `json:"incident,omitempty"`
}

func (o *StatusPageNotification) ToJson() string {
	b, _ := json.Marshal(o)
	return string(b)
}
