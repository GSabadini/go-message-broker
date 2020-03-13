package rabbitmq

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type Consumer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
}

func NewConsumer(conn *amqp.Connection, channel *amqp.Channel, queue amqp.Queue) Consumer {
	return Consumer{conn: conn, channel: channel, queue: queue}
}

func (c Consumer) Consume() error {
	deliveries, err := c.channel.Consume(
		c.queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("queue consume: %s", err)
	}

	//var done = make(chan error)
	//
	//go handle(deliveries, done)

	forever := make(chan bool)

	go func() {
		for d := range deliveries {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

	return nil
}

func handle(deliveries <-chan amqp.Delivery, done chan error) {
	for d := range deliveries {
		log.Printf(
			"got %dB delivery: [%v] %q",
			len(d.Body),
			d.DeliveryTag,
			d.Body,
		)
		d.Ack(false)
	}
	log.Printf("handle: deliveries channel closed")
	done <- nil
}
