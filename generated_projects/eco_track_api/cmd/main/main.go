package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"eco_track_api/internal/config"
	"eco_track_api/internal/controller"
	"eco_track_api/internal/model"
	"eco_track_api/internal/router"
	"eco_track_api/internal/service"
	"syscall"
	"time"
)

func main() {
	cfg, err := config.LoadConfig("./")
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot load config: %v\n", err)
		os.Exit(1)
	}

	db, err := model.ConnectDatabase(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot connect to database: %v\n", err)
		os.Exit(1)
	}

	carbonService := service.NewCarbonService(db)
	carbonController := controller.NewCarbonController(carbonService)

	r := router.NewRouter(&cfg, carbonController)
	r.SetRoutes()

	server := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: r.Gin,
	}

	go func() {
		log.Printf("Server is running on port %s\n", cfg.ServerPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server start failed: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown initiated...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server Shutdown Error: %v", err)
	}

	<-ctx.Done()
	log.Println("Server shutdown completed or timed out")

	log.Println("Server exiting")
	os.Exit(0)
}