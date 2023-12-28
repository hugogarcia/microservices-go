package event

import "github.com/rabbitmq/amqp091-go"

func declareExchange(channel *amqp091.Channel) error {
	return channel.ExchangeDeclare(
		"logs_topic", // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
}

func declareRandomQueue(channel *amqp091.Channel) (amqp091.Queue, error) {
	return channel.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
}
