package handlers

type Event struct {
	UserId      string `json:"userId"`
	WorkspaceId int64  `json:"workspaceId"`
	AnonymousId string `json:"anonymousId"`
	WriteKey    string `json:"writeKey"`
	EventType   string `json:"type"`
	Context     any    `json:"context"`
	ReceivedAt  string `json:"receivedAt"`
	SentAt      string `json:"sentAt"`
	Timestamp   string `json:"timestamp"`
}
