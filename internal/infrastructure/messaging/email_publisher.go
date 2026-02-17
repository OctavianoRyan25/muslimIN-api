package messaging

import (
	"encoding/json"

	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/domain"
	"github.com/rabbitmq/amqp091-go"
)

type EmailPublisherHandler interface {
	Publish(message domain.EmailMessage) error
}

type EmailPublisher struct {
	channel *amqp091.Channel
}

func NewEmailPublisher(ch *amqp091.Channel) EmailPublisherHandler {
	return &EmailPublisher{
		channel: ch,
	}
}

func (p *EmailPublisher) Publish(message domain.EmailMessage) error {

	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return p.channel.Publish(
		"",
		"email_queue",
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}
