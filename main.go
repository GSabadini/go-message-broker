package main

import (
	"fmt"
	"log"
	"os"

	"github.com/GSabadini/go-message-broker/kafka"
	"github.com/GSabadini/go-message-broker/rabbitmq"
	"github.com/Shopify/sarama"
)

const (
	PRODUCER = "producer"
	CONSUMER = "consumer"
	RABBITMQ = "rabbitmq"
	KAFKA = "kafka"
)

func main() {
	if os.Args[1] == RABBITMQ {
		startRabbitMQ()
	} else if os.Args[1] == KAFKA {
		startKafka()
	}
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

	var message = "Hello World RabbitMQ!"

	if os.Args[2] == PRODUCER {
		producer := rabbitmq.NewProducer(channel, queue.Name)
		if err := producer.Publish(message); err != nil {
			log.Fatalf("failed publish: %s", err)
		}
	} else if os.Args[2] == CONSUMER {
		consumer := rabbitmq.NewConsumer(channel, queue.Name)
		if err := consumer.Consume(); err != nil {
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

	if os.Args[2] == PRODUCER {
		kafka.NewProducer(config, logger, topic, brokers).Publish(message)
	} else if os.Args[2] == CONSUMER {
		kafka.NewConsumer(config, logger, topic, groupID, brokers).Consume()
	}
}

func startActiveMQ() {
	fmt.Println("implement me")
}

func startRedis() {
	fmt.Println("implement me")
}
