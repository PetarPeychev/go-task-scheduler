package config

import (
	"log"
	"os"

	"github.com/nats-io/nats.go"
)

type SchedulerConfig struct {
	NatsURL      string
	TasksSubject string
}

func LoadFromEnv() SchedulerConfig {
	natsURL := os.Getenv("NATS_URL")
	if natsURL == "" {
		natsURL = nats.DefaultURL
		log.Printf("NATS_URL environment variable is empty or not set, using default value: %s", natsURL)
	}

	tasksSubject := os.Getenv("TASKS_SUBJECT")
	if tasksSubject == "" {
		tasksSubject = "tasks"
		log.Printf("TASKS_SUBJECT environment variable is empty or not set, using default value: %s", tasksSubject)
	}

	return SchedulerConfig{
		NatsURL:      natsURL,
		TasksSubject: tasksSubject,
	}
}
