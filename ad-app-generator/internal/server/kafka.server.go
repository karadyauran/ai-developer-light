package server

import (
	"encoding/json"
	"fmt"
	"karadyaur.io/ai-dev-light/ad-app-generator/internal/generated"
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

	if err := c.consumer.Subscribe(c.config.KafkaTopic, nil); err != nil {
		return fmt.Errorf("failed to subscribe to Kafka topic: %w", err)
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	run := true

	for run {
		select {
		case sig := <-signals:
			fmt.Printf("Signal reached %v: finish work\n", sig)
			run = false
		default:
			msg, err := c.consumer.ReadMessage(100)
			if err != nil {
				if kafkaErr, ok := err.(kafka.Error); ok && kafkaErr.Code() == kafka.ErrTimedOut {
					continue
				}
				log.Printf("Error reading: %v", err)
				continue
			}

			var kafkareq generated.KafkaRequest
			if err := json.Unmarshal(msg.Value, &kafkareq); err != nil {
				log.Printf("Error marshaling: %v", err)
				continue
			}

			containerID, err := c.dockerService.StartContainer("internal/secret/Dockerfile")
			if err != nil {
				log.Printf("Faild to start container: %v", err)
				continue
			}

			log.Printf("Container started with ID: %s", containerID)
		}
	}

	return nil
}
