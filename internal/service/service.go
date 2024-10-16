package service

import "ai-dev-light/internal/config"

type Service struct {
	openAIService *OpenAIService
}

func NewService(config *config.Config) *Service {
	return &Service{
		openAIService: NewOpenAIService(config),
	}
}
