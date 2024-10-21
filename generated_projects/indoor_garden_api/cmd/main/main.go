package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"indoor_garden_api/internal/config"
	"indoor_garden_api/internal/controller"
	"indoor_garden_api/internal/router"
	"syscall"
	"time"
)

func main() {
	newConfig, err := config.LoadConfig("./")
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot load config: %v\n", err)
		os.Exit(1)
	}

	newController := controller.NewController()

	newRouter := router.NewRouter(&newConfig, newController)
	newRouter.SetRoutes()

	newServer := &http.Server{
		Addr:    ":" + newConfig.ServerPort,
		Handler: newRouter.Gin,
	}

	go func() {
		log.Printf("Server is running on port %s\n", newConfig.ServerPort)
		if err := newServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server start failed: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown initiated...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := newServer.Shutdown(ctx); err != nil {
		log.Printf("Server Shutdown Error: %v", err)
	}

	<-ctx.Done()
	log.Println("Server shutdown completed or timed out")

	log.Println("Server exiting")
	os.Exit(0)
}