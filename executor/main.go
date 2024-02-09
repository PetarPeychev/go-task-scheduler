package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	tasksQueueGroup := os.Getenv("TASKS_QUEUE_GROUP")
	if tasksQueueGroup == "" {
		log.Fatal("TASKS_QUEUE_GROUP environment variable is empty or not set")
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

	// Subscribe to the tasks subject with the specified queue group
	sub, err := nc.QueueSubscribe(tasksSubject, tasksQueueGroup, func(m *nats.Msg) {
		log.Printf("Received message on subject %q: %s", m.Subject, string(m.Data))
	})
	if err != nil {
		log.Fatalf("Failed to subscribe to subject %q: %v", tasksSubject, err)
	}
	defer sub.Unsubscribe()
	log.Printf("Subscribed to subject %q with queue group %q", tasksSubject, tasksQueueGroup)

	// Wait for an interrupt signal to gracefully shut down
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}
