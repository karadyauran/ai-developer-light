package server

import (
	"karadyaur.io/ai-dev-light/ad-app-generator/internal/config"
	"karadyaur.io/ai-dev-light/ad-app-generator/internal/service"
)

type Server struct {
	KafkaServer *KafkaServer
}

func NewServer(cfg *config.Config, service *service.Service) *Server {
	return &Server{
		KafkaServer: NewKafkaServer(cfg, service.DockerService),
	}
}
