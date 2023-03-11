package handlers

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAgentViewServer(t *testing.T) {
	var body = []byte(`{"userId": "some user", "writeKey": "some key"}`)
	var eventBody Event
	err := json.Unmarshal(body, &eventBody)
	assert.Nil(t, err)
	assert.Equal(t, "some user", eventBody.UserId)
}
