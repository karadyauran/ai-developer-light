package main

import (
	"karadyaur.io/ai-dev-light/ad-app-generator/internal/config"
	"karadyaur.io/ai-dev-light/ad-app-generator/internal/server"
	"karadyaur.io/ai-dev-light/ad-app-generator/internal/service"
	"log"
)

func main() {
	newConfig, err := config.LoadConfig("./")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	newService := service.NewService()
	newServer := server.NewServer(&newConfig, newService)

	if err := newServer.KafkaServer.Start(); err != nil {
		log.Fatalf("Ошибка при запуске KafkaHandler: %v", err)
	}
}
