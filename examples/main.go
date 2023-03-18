package main

import (
	"github.com/segmentio/analytics-go"
	"log"
	"time"
)

func main() {
	// Instantiates a client to use send messages to the segment API.
	client, _ := analytics.NewWithConfig(
		"some key",
		analytics.Config{
			Endpoint:  "http://localhost:5000",
			Interval:  60 * time.Second,
			BatchSize: 100,
			Verbose:   true,
		})

	// Enqueues a track event that will be sent asynchronously.
	if err := client.Enqueue(analytics.Track{
		Event:  "Download",
		UserId: "123456",
		Properties: map[string]any{
			"application": "Example App",
			"version":     "v1.1.0",
			"platform":    "Linux",
		},
	}); err != nil {
		log.Fatalf("error: %s", err)
	}

	// Flushes any queued messages and closes the client.
	if err := client.Close(); err != nil {
		log.Fatalf("error: %s", err)
	}
}
