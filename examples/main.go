package main

import (
	"github.com/segmentio/analytics-go"
	"log"
	"os"
)

func main() {
	// Instantiates a client to use send messages to the segment API.
	client := analytics.New(os.Getenv("some key"))

	// Enqueues a track event that will be sent asynchronously.
	if err := client.Enqueue(analytics.Track{
		UserId: "test-user",
		Event:  "test-snippet",
	}); err != nil {
		log.Fatalf("error: %s", err)
	}

	// Flushes any queued messages and closes the client.
	if err := client.Close(); err != nil {
		log.Fatalf("error: %s", err)
	}
}
