package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/infrastructure/mail"
	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/infrastructure/messaging"
)

func main() {
	// Context untuk graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Init RabbitMQ
	url := os.Getenv("RABBITMQ_URL")
	if url == "" {
		log.Fatal("RABBITMQ_URL is not set")
	}
	mq, err := messaging.NewRabbitMQ(url)
	if err != nil {
		log.Fatal("RabbitMQ connection failed:", err)
	}
	defer mq.Close()

	// Init SMTP
	mailing := mail.NewSmtpMailer(
		os.Getenv("SMTP_HOST"),
		os.Getenv("SMTP_PORT"),
		os.Getenv("SMTP_USER"),
		os.Getenv("SMTP_PASS"),
	)

	// Init Consumer
	consumer := messaging.NewEmailConsumer(mq.Channel, mailing)

	log.Println("Email Worker started...")

	// Start listening
	if err := consumer.Start(ctx); err != nil {
		log.Fatal("Consumer stopped:", err)
	}
}
