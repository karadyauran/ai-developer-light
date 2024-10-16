package service

import (
	"ai-dev-light/internal/config"
)

type Service struct {
	AppBuilder    *AppBuilder
	OpenAIService *OpenAIService
}

func NewService(cfg *config.Config, model string) (*Service, error) {
	openAIService := NewOpenAIService(cfg, model)
	appBuilder, err := NewAppBuilderService(openAIService)
	if err != nil {
		return nil, err
	}
	return &Service{
		OpenAIService: openAIService,
		AppBuilder:    appBuilder,
	}, nil
}
