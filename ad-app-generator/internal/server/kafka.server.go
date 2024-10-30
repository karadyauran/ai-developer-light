package server

import (
	"encoding/json"
	"fmt"
	"karadyaur.io/ai-dev-light/ad-app-generator/internal/generated"
	"karadyaur.io/ai-dev-light/ad-app-generator/internal/model"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"karadyaur.io/ai-dev-light/ad-app-generator/internal/config"
	"karadyaur.io/ai-dev-light/ad-app-generator/internal/service"
)

type KafkaServer struct {
	config        *config.Config
	dockerService *service.DockerService
	consumer      *kafka.Consumer
}

func NewKafkaServer(cfg *config.Config, dockerService *service.DockerService) *KafkaServer {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:" + cfg.KafkaServer,
		"group.id":          cfg.GroupID,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %v", err)
	}

	return &KafkaServer{
		config:        cfg,
		dockerService: dockerService,
		consumer:      consumer,
	}
}

func (c *KafkaServer) Start() error {
	defer c.consumer.Close()

	if err := c.consumer.Subscribe("app_generator", nil); err != nil {
		return fmt.Errorf("failed to subscribe to Kafka topic: %w", err)
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		default:
			msg, err := c.consumer.ReadMessage(-1)

			if err != nil {
				log.Fatal("Error received from Kafka server: ", err)
				return err
			}

			var kafkareq generated.KafkaRequest
			if err := json.Unmarshal(msg.Value, &kafkareq); err != nil {
				log.Fatalf("Error unmarshalling Kafka message: %v", err)
				return err
			}

			dockerContainer := model.DockerContainer{
				Hostname: kafkareq.Id,
				Image:    "ubuntu",
				Files: []model.File{
					{
						Filename: "internal/scripts/1-init.sh",
					},
					{
						Filename: "internal/scripts/2-setup.sh",
					},
					{
						Filename: "internal/scripts/3-generate_file.sh",
					},
					{
						Filename: "internal/scripts/4-cleanup.sh",
					},
				},
				Commands: []string{
					"chmod +x /app/1-init.sh",
					"chmod +x /app/2-setup.sh",
					"chmod +x /app/3-generate_file.sh",
					"chmod +x /app/4-cleanup.sh",
				},
			}

			run, err := c.dockerService.CopyAndRun(dockerContainer)
			if err != nil {
				return err
			}

			log.Println("Container ID" + run)
		}
	}
}
