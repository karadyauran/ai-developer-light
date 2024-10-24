package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"karadyaur.io/ai-dev-light/ad-oauth/internal/config"
	"karadyaur.io/ai-dev-light/ad-oauth/internal/generated"
	"karadyaur.io/ai-dev-light/ad-oauth/internal/server"
	"karadyaur.io/ai-dev-light/ad-oauth/internal/service"
	"karadyaur.io/ai-dev-light/ad-oauth/internal/utils"
	"karadyaur.io/ai-dev-light/ad-oauth/pkg/database"
	"log"
	"net"
	"os"
)

func main() {
	newConfig, err := config.LoadConfig("./")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	newPool, err := database.NewPostgresDB(&newConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer newPool.Close()

	lis, err := net.Listen("tcp", ":"+newConfig.GrpcAuthPort)
	if err != nil {
		log.Fatalf("failed to listen on gRPC port %s: %v", newConfig.GrpcAuthPort, err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	newGithubOAuth := utils.NewGitHubOAuth(&newConfig)
	newService := service.NewService(newPool, newGithubOAuth)
	newServer := server.NewAuthServer(newService)

	generated.RegisterUserServiceServer(grpcServer, newServer)

	// Start serving gRPC
	log.Printf("gRPC server listening on port %s", newConfig.GrpcAuthPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC: %v", err)
	}
}
