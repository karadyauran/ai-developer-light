package main

import (
	"fmt"
	"karadyaur.io/ai-dev-light/ad-app-generator/internal/config"
	"karadyaur.io/ai-dev-light/ad-app-generator/internal/model"
	"karadyaur.io/ai-dev-light/ad-app-generator/internal/service"
	"log"
)

func main() {
	_, err := config.LoadConfig("./")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	/*lis, err := net.Listen("tcp", ":"+newConfig.GrpcAIGeneratorPort)
	if err != nil {
		log.Fatalf("failed to listen on gRPC port %s: %v", newConfig.GrpcAIGeneratorPort, err)
	}*/

	newService := service.NewService()
	container, err := newService.DockerService.CreateContainer(model.DockerContainer{
		Hostname: "DockerName",
		Image:    "ubuntu",
	})
	if err != nil {
		log.Fatal("cannot create container:", err)
	}

	fmt.Printf("Container %s created successfully\n", container)
}
