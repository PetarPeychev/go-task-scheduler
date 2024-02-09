package main

import (
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
)

func main() {
	// Get configuration from environment variables
	natsURL := os.Getenv("NATS_URL")
	if natsURL == "" {
		log.Fatal("NATS_URL environment variable is empty or not set")
	}
	tasksSubject := os.Getenv("TASKS_SUBJECT")
	if tasksSubject == "" {
		log.Fatal("TASKS_SUBJECT environment variable is empty or not set")
	}

	// Connect to NATS
	nc, err := nats.Connect(
		natsURL,
		nats.RetryOnFailedConnect(true),
		nats.MaxReconnects(-1),
		nats.ReconnectWait(time.Second*5),
	)
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	defer nc.Close()
	log.Println("Connected to NATS")

	// Publish a message to the tasks subject every 10 seconds
	for {
		message := []byte(uuid.New().String())
		err := nc.Publish(tasksSubject, message)
		if err != nil {
			log.Printf("Failed to publish message: %v", err)
		} else {
			log.Printf("Published message to subject %q: %s", tasksSubject, message)
		}
		time.Sleep(time.Millisecond * 10) // Wait before sending the next message
	}
}
