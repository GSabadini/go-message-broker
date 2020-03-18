package main

import (
	"fmt"
	"log"
	"os"

	"github.com/GSabadini/go-message-broker/activemq"
	"github.com/GSabadini/go-message-broker/kafka"
	"github.com/GSabadini/go-message-broker/rabbitmq"

	"github.com/Shopify/sarama"
)

const (
	PRODUCER = "producer"
	CONSUMER = "consumer"

	RABBITMQ = "rabbitmq"
	KAFKA    = "kafka"
	ACTIVEMQ = "activemq"
)

func main() {
	switch os.Args[1] {
	case RABBITMQ:
		startRabbitMQ()
	case KAFKA:
		startKafka()
	case ACTIVEMQ:
		startActiveMQ()
	}
}

func startRabbitMQ() {
	connection, err := rabbitmq.OpenConnection()
	if err != nil {
		log.Fatalf("failed connection: %s", err)
	}
	defer func() {
		if err := connection.Close(); err != nil {
			log.Fatalf("failed close connection: %s", err)
		}
	}()

	channel, err := rabbitmq.NewChannel(connection).Create()
	if err != nil {
		log.Fatalf("failed create channel: %s", err)
	}

	queue, err := rabbitmq.NewQueue(channel, "go-message-broker").Create()
	if err != nil {
		log.Fatalf("failed queue declare: %s", err)
	}

	var message = "Hello World RabbitMQ!"

	switch os.Args[2] {
	case PRODUCER:
		if err := rabbitmq.NewProducer(channel, queue.Name).Publish(message); err != nil {
			log.Fatalf("failed publish message: %s", err)
		}
	case CONSUMER:
		if err := rabbitmq.NewConsumer(channel, queue.Name).Consume(); err != nil {
			log.Fatalf("failed consume: %s", err)
		}
	}
}

func startKafka() {
	var (
		config  = sarama.NewConfig()
		logger  = log.New(os.Stderr, "[sarama_logger]", log.LstdFlags)
		groupID = "sarama_consumer"
		topic   = "go-message-broker-topic"
		brokers = []string{"localhost:9093"}
		message = "Hello World Kafka!"
	)

	switch os.Args[2] {
	case PRODUCER:
		kafka.NewProducer(config, logger, topic, brokers).Publish(message)
	case CONSUMER:
		kafka.NewConsumer(config, logger, topic, groupID, brokers).Consume()
	}
}

func startActiveMQ() {
	connection, err := activemq.Connect()
	if err != nil {
		log.Fatalf("failed connection: %s", err)
	}
	defer func() {
		if err := connection.Disconnect(); err != nil {
			log.Fatalf("failed close connection: %s", err)
		}
	}()

	var (
		message = "Hello World ActiveMQ!"
		queue   = "go-message-broker"
	)

	switch os.Args[2] {
	case PRODUCER:
		if err := activemq.NewProducer(connection, queue).Publish(message); err != nil {
			log.Fatalf("failed publish message: %s", err)
		}
	case CONSUMER:
		activemq.NewConsumer(connection, queue).Consume()
	}
}

func startRedis() {
	fmt.Println("implement me")
}
