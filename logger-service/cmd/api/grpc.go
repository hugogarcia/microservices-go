package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/hugogarcia/microservices/logger-service/data"
	"github.com/hugogarcia/microservices/logger-service/logs"
	"google.golang.org/grpc"
)

type LogService struct {
	logs.UnimplementedLogServiceServer
	Models data.Models
}

func (l *LogService) WriteLog(ctx context.Context, in *logs.LogRequest) (*logs.LogResponse, error) {
	input := in.GetLogEntry()

	err := l.Models.LogEntry.Insert(ctx, data.LogEntry{
		Name: input.Name,
		Data: input.Data,
	})
	if err != nil {
		return &logs.LogResponse{
			Result: "failed to insert log",
		}, err
	}

	return &logs.LogResponse{
		Result: "success",
	}, nil
}

func (app *Config) grpcListen() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", gRpcPort))
	if err != nil {
		log.Fatal("Failed to listen gRPC server")
	}

	s := grpc.NewServer()
	logs.RegisterLogServiceServer(s, &LogService{Models: app.Models})
	log.Println("gRPC started on port: ", gRpcPort)

	if err := s.Serve(listener); err != nil {
		log.Fatal("Failed to start gRPC server")
	}
}
