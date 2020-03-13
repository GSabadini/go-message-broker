package main

import (
	"fmt"
	"log"
	"os"

	"github.com/GSabadini/go-message-broker/kafka"
	"github.com/GSabadini/go-message-broker/rabbitmq"
)

func main() {
	//startRabbitMQ()
	startKafka()
}

func startRabbitMQ() {
	connection, err := rabbitmq.OpenConnection()
	if err != nil {
		log.Fatalf("failed connection: %s", err)
	}
	defer connection.Close()

	channel, err := rabbitmq.NewChannel(connection).Create()
	if err != nil {
		log.Fatalf("failed create channel: %s", err)
	}

	queue, err := rabbitmq.NewQueue(channel).Create()
	if err != nil {
		log.Fatalf("failed queue declare: %s", err)
	}

	producer := rabbitmq.NewProducer(connection, channel, queue)
	if err := producer.Publish(); err != nil {
		log.Fatalf("failed publish: %s", err)
	}

	consumer := rabbitmq.NewConsumer(connection, channel, queue)
	if err := consumer.Consume(); err != nil {
		log.Fatalf("failed consume: %s", err)
	}
}

func startKafka() {
	if os.Args[1] == "producer" {
		kafka.Publish()
	} else if os.Args[1] == "consumer" {
		kafka.Consume()
	}
}

func startActiveMQ() {
	fmt.Println("implement me")
}
