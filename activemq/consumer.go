package activemq

import (
	"log"

	"github.com/go-stomp/stomp"
)

type Consumer struct {
	conn  *stomp.Conn
	queue string
}

func NewConsumer(conn *stomp.Conn, queue string) Consumer {
	return Consumer{conn: conn, queue: queue}
}

func (c Consumer) Consume() {
	sub, err := c.conn.Subscribe(c.queue, stomp.AckAuto)
	if err != nil {
		panic(err)
	}
	defer sub.Unsubscribe()

	for {
		m := <-sub.C
		log.Printf("Consumer received a message: %s in queue: %s", m.Body, c.queue)

		log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	}
}

func (c Consumer) handler(err error, message string) {
	if err != nil {
		log.Printf("Error consume message: %s", err)
	}

	log.Printf("Consumer received a message: %s", message)
}
