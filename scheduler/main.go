package main

import (
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	RABBITMQ_URL := os.Getenv("RABBITMQ_URL")

	// Connect to RabbitMQ
	var conn *amqp.Connection
	var err error
	for {
		conn, err = amqp.Dial(RABBITMQ_URL)
		if err != nil {
			log.Printf("Failed to connect to RabbitMQ: %v", err)
		} else {
			log.Println("Connected to RabbitMQ")
			break
		}
		time.Sleep(time.Second * 5)
	}
	defer conn.Close()

	for {
		time.Sleep(time.Second * 10)
		log.Println("Scheduler is running...")
	}

	// Continue with setting up channels, declaring queues, etc.
}
