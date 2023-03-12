package handlers

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
	WorkspaceId int64          `json:"workspaceId"`
	AnonymousId string         `json:"anonymousId"`
	MessageId   string         `json:"messageId"`
	WriteKey    string         `json:"writeKey"`
	EventType   string         `json:"type"`
	Event       string         `json:"event,omitempty"`
	Context     Context        `json:"context"`
	Properties  map[string]any `json:"properties"`
	ReceivedAt  string         `json:"receivedAt"`
	SentAt      string         `json:"sentAt"`
	Timestamp   string         `json:"timestamp"`
	Traits      Traits         `json:"traits,omitempty"`
}

type EventBatch struct {
	Event
	Batch []Event
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
