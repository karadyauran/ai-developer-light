package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"syscall"
	"wellness_api/internal/config"
	"wellness_api/internal/controller"
	"wellness_api/internal/router"
	"wellness_api/internal/database"
	"wellness_api/internal/service"
	"wellness_api/internal/repository"
)

func main() {
	newConfig, err := config.LoadConfig("./")
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot load config: %v\n", err)
		os.Exit(1)
	}

	db, err := database.Connect(newConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	mainRouter := router.NewRouter(&newConfig, userController)
	mainRouter.SetRoutes()

	server := &http.Server{
		Addr:    ":" + newConfig.ServerPort,
		Handler: mainRouter.Gin,
	}

	go func() {
		log.Printf("Server is running on port %s\n", newConfig.ServerPort)
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
}