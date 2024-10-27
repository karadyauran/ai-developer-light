package service

import (
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"karadyaur.io/ai-dev-light/ad-kafka/internal/model"
	"log"
)

type KafkaProducer struct {
	Producer *kafka.Producer
}

func NewKafkaProducer(brokers string) *KafkaProducer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": brokers})
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return &KafkaProducer{
		Producer: p,
	}
}

func (kp *KafkaProducer) Send(request model.Request) error {
	messageBytes, err := json.Marshal(request)
	if err != nil {
		log.Printf("Failed to marshal request: %v", err)
		return err
	}

	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &request.Topic, Partition: kafka.PartitionAny},
		Value:          messageBytes,
	}

	err = kp.Producer.Produce(msg, nil)
	if err != nil {
		log.Printf("Failed to send request: %v", err)
		return err
	}

	return nil
}
