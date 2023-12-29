package main

import (
	"fmt"
	"log"
	"time"

	"github.com/hugogarcia/microservices/listener-services/event"
	"github.com/rabbitmq/amqp091-go"
)

func main() {
	rabbitmq, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	defer rabbitmq.Close()
	fmt.Println("Connected to RabbitMQ")

	log.Println("listening and consuming RabbitMQ")
	consumer, err := event.NewConsumer(rabbitmq)
	if err != nil {
		log.Fatal(err)
	}

	err = consumer.Listen([]string{"log.INFO", "log.WARNING", "log.ERROR"})

}

func connect() (*amqp091.Connection, error) {
	var counts int64
	var backoff = 2 * time.Second

	for {
		connection, err := amqp091.Dial("amqp://guest:guest@rabbitmq")
		if err == nil {
			return connection, nil
		}

		counts++
		if counts > 5 {
			fmt.Println(err)
			return nil, err
		}

		fmt.Println("RabbitMQ not ready, trying again...")
		time.Sleep(backoff)
	}
}
