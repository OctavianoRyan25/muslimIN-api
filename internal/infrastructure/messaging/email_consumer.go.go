package messaging

import (
	"context"
	"encoding/json"
	"log"

	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/infrastructure/mail"
	"github.com/rabbitmq/amqp091-go"
)

type EmailPayload struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type EmailConsumer struct {
	channel *amqp091.Channel
	mail    mail.UserMail
}

func NewEmailConsumer(ch *amqp091.Channel, mail mail.UserMail) *EmailConsumer {
	return &EmailConsumer{
		channel: ch,
		mail:    mail,
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
			err := c.mail.SendEmailToUser([]string{payload.To}, payload.Subject, payload.Body)
			if err != nil {
				log.Printf("Gagal kirim email: %v", err)
				msg.Nack(false, true)
				continue
			} else {
				log.Printf("Success sending email to %s", payload.To)
				msg.Ack(false) // Tandai sukses
			}
		}
	}
}
