package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"generated_projects/carbon_track/internal/config"
	"generated_projects/carbon_track/internal/controller"
	"generated_projects/carbon_track/internal/model"
	"generated_projects/carbon_track/internal/repository"
	"generated_projects/carbon_track/internal/router"
	"generated_projects/carbon_track/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig("./")
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot load config: %v\n", err)
		os.Exit(1)
	}

	db, err := model.ConnectPostgres(cfg)
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	repo := repository.NewCarbonRepository(db)
	svc := service.NewCarbonService(repo)
	ctrl := controller.NewCarbonController(svc)

	r := gin.Default()
	router.SetupRoutes(r, ctrl)

	srv := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: r,
	}

	go func() {
		log.Printf("Server is running on port %s\n", cfg.ServerPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown initiated...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown:", err)
	}

	<-ctx.Done()
	log.Println("Server exiting")
}
