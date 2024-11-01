package server

import (
	"context"
	"karadyaur.io/ai-dev-light/ad-kafka/internal/generated"
	"karadyaur.io/ai-dev-light/ad-kafka/internal/model"
	"karadyaur.io/ai-dev-light/ad-kafka/internal/service"
)

type KafkaServer struct {
	generated.UnimplementedKafkaServiceServer
	kafkaService *service.KafkaProducer
}

func NewKafkaServer(kafkaService *service.Service) *KafkaServer {
	return &KafkaServer{
		kafkaService: kafkaService.KafkaProducer,
	}
}

func (ks *KafkaServer) KafkaSend(ctx context.Context, request *generated.KafkaRequest) (*generated.KafkaResponse, error) {
	sendMessageRequest := model.Request{
		ID:    request.Id,
		Topic: request.Topic,
		Params: model.Params{
			ProjectType:    request.Params.ProjectType,
			Language:       request.Params.Language,
			GenerationType: request.Params.GenerationType,
			Idea:           request.Params.Idea,
			Token:          request.Params.Token,
		},
		CreatedAt: request.CreatedAt,
	}

	err := ks.kafkaService.Send(sendMessageRequest)
	if err != nil {
		return &generated.KafkaResponse{
			Response: "Error creating message",
		}, err
	}

	return &generated.KafkaResponse{Response: "Request sent successfully"}, nil
}
