package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"karadyaur.io/ai-dev-light/ad-kafka/internal/config"
	"karadyaur.io/ai-dev-light/ad-kafka/internal/generated"
	"karadyaur.io/ai-dev-light/ad-kafka/internal/server"
	"karadyaur.io/ai-dev-light/ad-kafka/internal/service"
	"log"
	"net"
)

func main() {
	newConfig, err := config.LoadKafkaConfig("./")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	lis, err := net.Listen("tcp", ":"+newConfig.ServerPort)
	if err != nil {
		log.Fatalf("failed to listen on gRPC port %s: %v", newConfig.ServerPort, err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	newService := service.NewService(newConfig.Brokers)
	newServer := server.NewKafkaServer(newService)

	generated.RegisterKafkaServiceServer(grpcServer, newServer)

	// Start serving gRPC
	log.Printf("gRPC server listening on port %s", newConfig.ServerPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC: %v", err)
	}
}
