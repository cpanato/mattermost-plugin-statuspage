package main

import (
	"encoding/json"
	"time"
)

type Meta struct {
	GeneratedAt   time.Time `json:"generated_at,omitempty"`
	Unsubscribe   string    `json:"unsubscribe,omitempty"`
	Documentation string    `json:"documentation,omitempty"`
}

type Page struct {
	ID                string `json:"id,omitempty"`
	StatusIndicator   string `json:"status_indicator,omitempty"`
	StatusDescription string `json:"status_description,omitempty"`
}

type Component struct {
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	Description string    `json:"description,omitempty"`
	ID          string    `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	PageID      string    `json:"page_id,omitempty"`
	Status      string    `json:"status,omitempty"`
	GroupID     string    `json:"group_id,omitempty"`
	Position    int       `json:"position,omitempty"`
	Showcase    bool      `json:"showcase,omitempty"`
}

type ComponentUpdate struct {
	CreatedAt     time.Time `json:"created_at,omitempty"`
	OldStatus     string    `json:"old_status,omitempty"`
	NewStatus     string    `json:"new_status,omitempty"`
	ComponentType string    `json:"component_type,omitempty"`
	State         string    `json:"state,omitempty"`
	ID            string    `json:"id,omitempty"`
	ComponentID   string    `json:"component_id,omitempty"`
}

type IncidentUpdate struct {
	CreatedAt          time.Time `json:"created_at,omitempty"`
	DisplayAt          time.Time `json:"display_at,omitempty"`
	TwitterUpdatedAt   time.Time `json:"twitter_updated_at,omitempty"`
	UpdatedAt          time.Time `json:"updated_at,omitempty"`
	ID                 string    `json:"id,omitempty"`
	Body               string    `json:"body,omitempty"`
	IncidentID         string    `json:"incident_id,omitempty"`
	Status             string    `json:"status,omitempty"`
	WantsTwitterUpdate bool      `json:"wants_twitter-update,omitempty"`
}

type Incident struct { // nolint: govet
	IncidentUpdates               []*IncidentUpdate `json:"incident_updates,omitempty"`
	Components                    []*Component      `json:"components,omitempty"`
	CreatedAt                     time.Time         `json:"created_at,omitempty"`
	MonitoringAt                  time.Time         `json:"monitoring_at,omitempty"`
	PostmortemBodyLastUpdatedAt   time.Time         `json:"postmortem_body_last_updated_at,omitempty"`
	PostmortemPublishedAt         time.Time         `json:"postmorem_published_at,omitempty"`
	ResolvedAt                    time.Time         `json:"resolved_at,omitempty"`
	ScheduledFor                  time.Time         `json:"scheduled_for"`
	ScheduledRemindedAt           time.Time         `json:"scheduled_reminded_at,omitempty"`
	ScheduledUntil                time.Time         `json:"scheduled_until,omitempty"`
	UpdatedAt                     time.Time         `json:"updated_at,omitempty"`
	ID                            string            `json:"id,omitempty"`
	Impact                        string            `json:"impact,omitempty"`
	ImpactOverride                string            `json:"impact_override,omitempty"`
	Name                          string            `json:"name,omitempty"`
	PageID                        string            `json:"page_id,omitempty"`
	PostmortemBody                string            `json:"postmortem_body,omitempty"`
	Shortlink                     string            `json:"shortlink,omitempty"`
	Status                        string            `json:"status,omitempty"`
	Backfilled                    bool              `json:"backfilled,omitempty"`
	PostmortemIgnored             bool              `json:"postmortem_ignored,omitempty"`
	PostmortemNotifiedSubscribers bool              `json:"postmortem_notified_subscribers,omitempty"`
	PostmortemNotifiedTwitter     bool              `json:"postmortem_notified_twitter,omitempty"`
	ScheduledAutoInProgress       bool              `json:"scheduled_auto_in_progress,omitempty"`
	ScheduledAutoCompleted        bool              `json:"scheduled_auto_completed,omitempty"`
	ScheduledRemindPrior          bool              `json:"scheduled_remind_prior,omitempty"`
}

type StatusPageNotification struct {
	Meta            *Meta            `json:"meta,omitempty"`
	Page            *Page            `json:"page,omitempty"`
	Component       *Component       `json:"component,omitempty"`
	ComponentUpdate *ComponentUpdate `json:"component_update,omitempty"`
	Incident        *Incident        `json:"incident,omitempty"`
}

func (o *StatusPageNotification) ToJSON() string {
	b, _ := json.Marshal(o)
	return string(b)
}
