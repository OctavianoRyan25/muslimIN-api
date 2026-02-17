package messaging

import (
	"context"
	"encoding/json"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type EmailPayload struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type EmailConsumer struct {
	channel *amqp091.Channel
}

func NewEmailConsumer(ch *amqp091.Channel) *EmailConsumer {
	return &EmailConsumer{
		channel: ch,
	}
}

func (c *EmailConsumer) Start(ctx context.Context) error {
	msgs, err := c.channel.Consume(
		"email_queue",
		"",    // consumer name
		false, // auto-ack false (agar bisa retry)
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	log.Println("[EmailConsumer] Waiting for messages...")

	for {
		select {
		case <-ctx.Done():
			log.Println("[EmailConsumer] Shutting down...")
			return nil

		case msg := <-msgs:
			var payload EmailPayload

			if err := json.Unmarshal(msg.Body, &payload); err != nil {
				log.Println("Invalid email payload:", err)
				msg.Nack(false, false) // reject tanpa retry
				continue
			}

			// Simulasi kirim email (sementara pakai log)
			log.Printf("[EmailConsumer] Send Email -> To: %s | Subject: %s",
				payload.To, payload.Subject)

			log.Printf("[EmailConsumer] Body: %s", payload.Body)

			// Ack = sukses diproses
			msg.Ack(false)
		}
	}
}
