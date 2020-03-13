package kafka

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/Shopify/sarama"
)

type Producer struct {
	conn sarama.AsyncProducer
}

func NewProducer(kafkaBrokers []string, config *sarama.Config) (*Producer, error) {

	conn, err := sarama.NewAsyncProducer(kafkaBrokers, config)
	if err != nil {
		return nil, errors.New("failed")
	}

	return &Producer{conn: conn }, nil
}

var (
	kafkaBrokers = []string{"localhost:9093"}
	KafkaTopic   = "sarama_topic"
	enqueued     int
)

func Publish() {
	producer, err := setupProducer()
	if err != nil {
		panic(err)
	} else {
		log.Println("Kafka AsyncProducer up and running!")
	}

	// Trap SIGINT to trigger a graceful shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	produceMessages(producer, signals)

	log.Printf("Kafka AsyncProducer finished with %d messages produced.", enqueued)
}

// setupProducer will create a AsyncProducer and returns it
func setupProducer() (sarama.AsyncProducer, error) {
	config := sarama.NewConfig()
	sarama.Logger = log.New(os.Stderr, "[sarama_logger]", log.LstdFlags)

	return sarama.NewAsyncProducer(kafkaBrokers, config)
}

// produceMessages will send 'Hello World Kafka!' to KafkaTopic each second, until receive a os signal to stop e.g. control + c
// by the user in terminal
func produceMessages(producer sarama.AsyncProducer, signals chan os.Signal) {
	for {
		time.Sleep(time.Second)
		message := &sarama.ProducerMessage{Topic: KafkaTopic, Value: sarama.StringEncoder("Hello World Kafka!")}
		
		select {
		case producer.Input() <- message:
			enqueued++
			log.Println("New Message produced")
		case <-signals:
			producer.AsyncClose() // Trigger a shutdown of the producer.
			return
		}
	}
}
