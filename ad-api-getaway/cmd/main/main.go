package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"karadyaur.io/ai-dev-light/ad-api-getaway/internal/config"
	"karadyaur.io/ai-dev-light/ad-api-getaway/internal/controller"
	"karadyaur.io/ai-dev-light/ad-api-getaway/internal/generated"
	routers "karadyaur.io/ai-dev-light/ad-api-getaway/internal/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title           AI Dev entry point
// @version         1.0
// @description     This API handles all requests from the users to microservices

// @license.name    MIT
// @license.url     https://github.com/karadyauran/ai-developer-light/blob/main/LICENSE

// @host            localhost:8000
// @BasePath        /api/v1
func main() {
	newConfig, err := config.LoadConfig("./")
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot load config: %v\n", err)
		os.Exit(1)
	}

	connOAuth, err := grpc.NewClient("localhost"+":"+newConfig.GrpcOAuthPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to Auth Service: %v", err)
	}
	defer connOAuth.Close()
	oAuthServiceClient := generated.NewOAuthServiceClient(connOAuth)

	connKafka, err := grpc.NewClient("localhost"+":"+newConfig.GrpcKafka, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to Kafka Service: %v", err)
	}
	defer connOAuth.Close()
	kafkaServiceClient := generated.NewKafkaServiceClient(connKafka)

	newController := controller.NewController(oAuthServiceClient, kafkaServiceClient)

	newRouter := routers.NewRouter(&newConfig, newController)
	newRouter.SetRoutes()

	newServer := &http.Server{
		Addr:    ":" + newConfig.ServerPort,
		Handler: newRouter.Gin,
	}

	// Start the server in a separate goroutine
	go func() {
		log.Printf("Server is running on port %s\n", newConfig.ServerPort)
		if err := newServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server start failed: %s\n", err)
		}
	}()

	// Set up signal catching
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown initiated...")

	// Context for graceful shutdown with a timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to gracefully shutdown the server
	if err := newServer.Shutdown(ctx); err != nil {
		log.Printf("Server Shutdown Error: %v", err)
	}

	// Waiting for the shutdown context to be done or timeout
	<-ctx.Done()
	log.Println("Server shutdown completed or timed out")

	log.Println("Server exiting")
	os.Exit(0)
}
