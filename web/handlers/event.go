package handlers

type Event struct {
	UserId      string         `json:"userId"`
	WorkspaceId int64          `json:"workspaceId"`
	AnonymousId string         `json:"anonymousId"`
	WriteKey    string         `json:"writeKey"`
	EventType   string         `json:"type"`
	Event       string         `json:"event"`
	Context     map[string]any `json:"context"`
	Properties  map[string]any `json:"properties"`
	ReceivedAt  string         `json:"receivedAt"`
	SentAt      string         `json:"sentAt"`
	Timestamp   string         `json:"timestamp"`
}

type EventBatch struct {
	Event
	Batch []Event
}

func (e *Event) addContextProperties(event *Event) {
	for property, value := range event.Context {
		e.Context[property] = value
	}
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
