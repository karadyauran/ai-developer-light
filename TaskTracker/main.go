package main

import (
    "log"
    "net/http"
    "project/config"
    "project/routes"
)

func main() {
    // Load configuration
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Error loading config: %v", err)
    }

    // Initialize routes
    router := routes.InitRoutes()

    // Start server
    log.Printf("Starting server on %s", cfg.ServerAddress)
    if err := http.ListenAndServe(cfg.ServerAddress, router); err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
}