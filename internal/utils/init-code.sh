#!/bin/bash

# Get the directory where the script is located
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Set the root directory to the parent of the script directory
ROOT_DIRECTORY="$(dirname "$(dirname "$SCRIPT_DIR")")"

# Get the project name from the first argument
PROJECT_NAME="$1"

# Check if the project name is provided
if [ -z "$PROJECT_NAME" ]; then
    echo "No project name provided."
    exit 1
fi

# Define the generated projects directory and project directory
GENERATED_PROJECTS_DIR="$ROOT_DIRECTORY/generated_projects"
PROJECT_DIRECTORY="$GENERATED_PROJECTS_DIR/$PROJECT_NAME"




# FILL FILES

# Define the path
CONFIG_DIR="$PROJECT_DIRECTORY/internal/config"
CONFIG_FILE="$CONFIG_DIR/config.go"

# Create the directory if it doesn't exist
mkdir -p "$CONFIG_DIR"

# Write the Go code to the config.go file
cat > "$CONFIG_FILE" << 'EOF'
package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	EnvType       string `mapstructure:"ENV_TYPE"`
	ServerPort    string `mapstructure:"SERVER_PORT"`
	WebappBaseUrl string `mapstructure:"WEBAPP_BASE_URL"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.SetConfigFile(path + ".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("could not load config: %v", err)
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("could not unmarshal config: %v", err)
	}

	return
}
EOF

# Define the path
MAIN_DIR="$PROJECT_DIRECTORY/cmd/main"
MAIN_FILE="$MAIN_DIR/main.go"

# Create the directory if it doesn't exist
mkdir -p "$MAIN_DIR"

# Write the Go code
cat > "$MAIN_FILE" << EOF
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"generated_projects/$PROJECT_NAME/internal/config"
	"generated_projects/$PROJECT_NAME/internal/controllers"
	"generated_projects/$PROJECT_NAME/internal/generated"
	"generated_projects/$PROJECT_NAME/internal/routers"
	"syscall"
	"time"
)

func main() {
	newConfig, err := config.LoadConfig("./")
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot load config: %v\n", err)
		os.Exit(1)
	}

	newController := controllers.NewController(authServiceClient)

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
EOF

# Define the path where the config.go file should be created
CONTROLLER_DIR="$PROJECT_DIRECTORY/internal/controller"
CONTROLLER_FILE="$CONTROLLER_DIR/controller.go"

# Create the directory if it doesn't exist
mkdir -p "$CONTROLLER_DIR"

# Write the Go code to the config.go file
cat > "$CONTROLLER_FILE" << EOF
package controller

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}
EOF

# Define the path where the config.go file should be created
ROUTER_DIR="$PROJECT_DIRECTORY/internal/router"
ROUTER_FILE="$ROUTER_DIR/router.go"

# Create the directory if it doesn't exist
mkdir -p "$ROUTER_DIR"

# Write the Go code to the config.go file
cat > "$ROUTER_FILE" << EOF
package router

type Router struct {
}

func NewRouter() *Router {
	return &Router{}
}
EOF

# Define the path where the config.go file should be created
SERVICE_DIR="$PROJECT_DIRECTORY/internal/service"
SERVICE_FILE="$SERVICE_DIR/service.go"

# Create the directory if it doesn't exist
mkdir -p "$SERVICE_DIR"

# Write the Go code to the config.go file
cat > "$SERVICE_FILE" << EOF
package router

type Service struct {
}

func NewService() *Service {
	return &Service{}
}
EOF

go mod init "$PROJECT_NAME"
mv go.mod "$GENERATED_PROJECTS_DIR/$PROJECT_NAME"

echo "Files created."