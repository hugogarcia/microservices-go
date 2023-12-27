package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hugogarcia/microservices/logger-service/data"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	webPort  = "8282"
	rpcPort  = "5001"
	mongoURL = "mongodb://mongo:27017"
	gRpcPort = "50001"
)

var client *mongo.Client

type Config struct {
	Models data.Models
}

func main() {
	mongoClient, err := connectToMongo()
	if err != nil {
		log.Panic(err)
	}

	client = mongoClient

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	defer client.Disconnect(ctx)

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Panic(err)
		}
	}()

	app := Config{
		Models: data.New(client),
	}

	app.server()
}

func (app *Config) server() {
	server := http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	log.Println(fmt.Sprintf("service started at port %s", webPort))
	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func connectToMongo() (*mongo.Client, error) {
	mongoOpts := options.Client().
		ApplyURI(mongoURL).
		SetAuth(options.Credential{
			Username:   "admin",
			Password:   "password",
			AuthSource: "admin",
		})

	ctx := context.Background()
	count := 0
	var err error
	for count < 3 {
		time.Sleep(time.Second)

		client, err := mongo.Connect(ctx, mongoOpts)
		if err != nil {
			count++
			log.Println("Error connecting to MongoDB")
			continue
		}

		if err = client.Ping(ctx, readpref.Primary()); err != nil {
			count++
			log.Println("Error pinging MongoDB")
			continue
		}
		log.Println("Connected to MongoDB")
		return client, nil
	}
	fmt.Println(err)
	return nil, err
}
