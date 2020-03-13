package kafka

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Shopify/sarama"
)

//type Consumer struct {
//	conn kafka.Consumer
//}

var (
	//kafkaBrokers = []string{"localhost:9093"}
	kafkaTopics     = []string{"sarama_topic"}
	consumerGroupID = "sarama_consumer"
)

func Consume() {
	// Init config, specify appropriate version
	config := sarama.NewConfig()
	sarama.Logger = log.New(os.Stderr, "[sarama_logger]", log.LstdFlags)
	config.Version = sarama.V2_1_0_0

	// Start with a client
	client, err := sarama.NewClient(kafkaBrokers, config)
	if err != nil {
		panic(err)
	}
	defer func() { _ = client.Close() }()

	// Start a new consumer group
	group, err := sarama.NewConsumerGroupFromClient(consumerGroupID, client)
	if err != nil {
		panic(err)
	}
	defer func() { _ = group.Close() }()
	log.Println("Consumer up and running")

	// Track errors
	go func() {
		for err := range group.Errors() {
			fmt.Println("ERROR", err)
		}
	}()

	// Iterate over consumer sessions.
	ctx := context.Background()
	for {
		handler := ConsumerGroupHandler{}

		err := group.Consume(ctx, kafkaTopics, handler)
		if err != nil {
			panic(err)
		}
	}
}
