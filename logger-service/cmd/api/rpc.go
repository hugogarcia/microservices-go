package main

import (
	"context"
	"log"
	"time"

	"github.com/hugogarcia/microservices/logger-service/data"
)

type RPCServer struct {
}

type RPCPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (r *RPCServer) LogInfo(payload *RPCPayload, resp *string) error {
	collection := client.Database("logs").Collection("logs")
	_, err := collection.InsertOne(context.TODO(), data.LogEntry{
		Name:      payload.Name,
		Data:      payload.Data,
		CreatedAt: time.Now().String(),
	})
	if err != nil {
		log.Println("error inserting mongo")
		return err
	}

	*resp = "Processed payload via RPC:" + payload.Name

	return nil
}
