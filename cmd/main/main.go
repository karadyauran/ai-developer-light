package main

import (
	"ai-dev-light/internal/config"
	"ai-dev-light/internal/service"
	"fmt"
	"os"
)

func main() {
	newConfig, err := config.LoadConfig("./")
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "cannot load config: %v\n", err)
		if err != nil {
			return
		}
		os.Exit(1)
	}

	newService := service.NewService(&newConfig)
	err = newService.AppBuilder.BuildWithNoContext()
	if err != nil {
		return
	}

	/*newController := controller.NewController(newService)
	newRouter := routers.NewRouter(&newConfig, newController)

	newServer := &http.Server{
		Addr:    ":" + newConfig.ServerPort,
		Handler: newRouter.Gin,
	}

	// Start the server in a separate goroutine
	go func() {
		log.Printf("Server is running on port %s\n", newConfig.ServerPort)
		if err := newServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
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

	// Attempt to gracefully shut down the server
	if err := newServer.Shutdown(ctx); err != nil {
		log.Printf("Server Shutdown Error: %v", err)
	}

	// Waiting for the shutdown context to be done or timeout
	<-ctx.Done()
	log.Println("Server shutdown completed or timed out")

	log.Println("Server exiting")
	os.Exit(0)*/
}
