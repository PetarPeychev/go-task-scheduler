package main

import (
	"log"
	"time"

	"github.com/PetarPeychev/go-task-scheduler/scheduler/config"
	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
)

func main() {
	config := config.LoadFromEnv()

	// Connect to NATS
	nc, err := nats.Connect(
		config.NatsURL,
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
		err := nc.Publish(config.TasksSubject, message)
		if err != nil {
			log.Printf("Failed to publish message: %v", err)
		} else {
			log.Printf("Published message to subject %q: %s", config.TasksSubject, message)
		}
		time.Sleep(time.Millisecond * 10) // Wait before sending the next message
	}
}
