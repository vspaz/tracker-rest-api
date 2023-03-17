package handlers

import "time"

type ContextLibrary struct {
	Group   string `json:"group"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

type ContextPage struct {
	Referrer string `json:"referrer"`
	InIframe bool   `json:"inIframe"`
}

type Context struct {
	Library   ContextLibrary `json:"library"`
	Page      ContextPage    `json:"page"`
	UserAgent string         `json:"userAgent"`
}

type Traits struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Event struct {
	UserId      string         `json:"userId"`
	SentAt      string         `json:"sentAt"`
	WorkspaceId int64          `json:"workspaceId"`
	AnonymousId string         `json:"anonymousId"`
	MessageId   string         `json:"messageId"`
	WriteKey    string         `json:"writeKey"`
	EventType   string         `json:"type"`
	Event       string         `json:"event,omitempty"`
	Context     Context        `json:"context"`
	Properties  map[string]any `json:"properties"`
	ReceivedAt  string         `json:"receivedAt"`
	Timestamp   string         `json:"timestamp"`
	Traits      Traits         `json:"traits,omitempty"`
}

type EventBatch struct {
	MessageId string    `json:"messageId"`
	SentAt    time.Time `json:"sentAt"`
	Messages  []Event   `json:"batch"`
	Context   *Context  `json:"context"`
}

func (e *Event) addContextProperties(event *Event) {
	e.Context = event.Context
}

func (e *Event) addUserDefinedProperties(event *Event) {
	for property, value := range event.Properties {
		e.Properties[property] = value
	}
}

func (e *Event) Merge(event *Event) {
	if event.AnonymousId != "" {
		e.AnonymousId = event.AnonymousId
	}

	if event.UserId != "" {
		e.UserId = event.UserId
	}

	if event.WorkspaceId != 0 {
		e.WorkspaceId = event.WorkspaceId
	}

	if event.WriteKey != "" {
		e.WriteKey = event.WriteKey
	}

	if event.Timestamp != "" {
		e.Timestamp = event.Timestamp
	}

	if event.SentAt != "" {
		e.SentAt = event.SentAt
	}
	e.addContextProperties(event)
}
