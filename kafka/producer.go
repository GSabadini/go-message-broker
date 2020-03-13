package kafka

import (
	kafka "github.com/Shopify/sarama"
)

type Producer struct {
	conn kafka.AsyncProducer
}
