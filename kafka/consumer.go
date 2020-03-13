package kafka

import (
	kafka "github.com/Shopify/sarama"
)

type Consumer struct {
	conn kafka.Consumer
}
