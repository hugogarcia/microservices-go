package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

const webPort = "8080"

type Config struct {
	Rabbit *amqp091.Connection
}

func main() {
	log.Printf("Starting broker on port %s\n", webPort)

	conn, err := connectToRabbit()
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Println("Connected to RabbitMQ")

	app := Config{
		Rabbit: conn,
	}
	server := http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func connectToRabbit() (*amqp091.Connection, error) {
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
