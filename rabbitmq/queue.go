package rabbitmq

import "github.com/streadway/amqp"

type Queue struct {
	channel *amqp.Channel
}

func NewQueue(channel *amqp.Channel) Queue {
	return Queue{channel: channel}
}

func (q Queue) Create() (amqp.Queue, error) {
	return q.channel.QueueDeclare(
		"go-message-broker",
		false,
		false,
		false,
		false,
		nil,
	)
}
