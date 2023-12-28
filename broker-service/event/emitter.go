package event

import (
	"context"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type Emitter struct {
	conn *amqp091.Connection
}

func (e *Emitter) setup() error {
	channel, err := e.conn.Channel()
	if err != nil {
		return err
	}

	return declareExchange(channel)
}

func (e *Emitter) Push(event string, severity string) error {
	channel, err := e.conn.Channel()
	if err != nil {
		return err
	}

	defer channel.Close()
	log.Println("Pushing to channel")

	err = channel.PublishWithContext(
		context.Background(),
		"logs_topic",
		severity,
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(event),
		},
	)

	if err != nil {
		return err
	}

	return nil
}
func NewEventEmitter(conn *amqp091.Connection) (Emitter, error) {
	emitter := Emitter{
		conn: conn,
	}

	err := emitter.setup()
	if err != nil {
		return Emitter{}, err
	}

	return emitter, nil
}
