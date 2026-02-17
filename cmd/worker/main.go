package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/infrastructure/messaging"
)

func main() {
	// Context untuk graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Init RabbitMQ
	mq, err := messaging.NewRabbitMQ("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("RabbitMQ connection failed:", err)
	}
	defer mq.Close()

	// Init Consumer
	consumer := messaging.NewEmailConsumer(mq.Channel)

	log.Println("Email Worker started...")

	// Start listening
	if err := consumer.Start(ctx); err != nil {
		log.Fatal("Consumer stopped:", err)
	}
}
