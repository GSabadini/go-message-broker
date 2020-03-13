package rabbitmq

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type Producer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
}

func NewProducer(conn *amqp.Connection, channel *amqp.Channel, queue amqp.Queue) Producer {
	return Producer{conn: conn, channel: channel, queue: queue}
}

func (p Producer) Publish(message string) error {
	if err := p.channel.Publish(
		"",
		p.queue.Name,
		false,
		false,
		amqp.Publishing{
			Headers:     amqp.Table{},
			ContentType: "text/plain",
			Body:        []byte(message),
		}); err != nil {
		return fmt.Errorf("failed to publish a message: %s", err)
	}

	log.Printf("New message publish:  %s", message)

	return nil
}
