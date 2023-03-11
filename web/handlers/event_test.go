package handlers

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEventDeserializationOk(t *testing.T) {
	var body = []byte(`{"sentAt":"2023-03-11T19:23:48.022Z","timestamp":"2023-03-11T19:23:42.877Z","event":"something_happened","type":"track","userId":"12345"}`)
	var eventBody Event
	err := json.Unmarshal(body, &eventBody)
	assert.Nil(t, err)
	assert.Equal(t, "2023-03-11T19:23:48.022Z", eventBody.SentAt)
	assert.Equal(t, "2023-03-11T19:23:42.877Z", eventBody.Timestamp)
	assert.Equal(t, "something_happened", eventBody.Event)
	assert.Equal(t, "track", eventBody.EventType)
	assert.Equal(t, "12345", eventBody.UserId)
}
