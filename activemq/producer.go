package activemq

import (
	"fmt"
	"log"

	"github.com/go-stomp/stomp"
)

type Producer struct {
	conn  *stomp.Conn
	queue string
}

func NewProducer(conn *stomp.Conn, queue string) Producer {
	return Producer{conn: conn, queue: queue}
}

func (p Producer) Publish(message string) error {
	if err := p.conn.Send(
		p.queue,
		"text/plain",
		[]byte(message),
	); err != nil {
		return fmt.Errorf("failed to publish a message: %s", err)
	}

	log.Printf("New message publish:  %s", message)

	return nil
}
