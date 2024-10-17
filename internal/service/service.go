package service

import (
	"ai-dev-light/internal/config"
)

type Service struct {
	AppBuilder *AppBuilder
}

func NewService(config *config.Config) *Service {
	return &Service{
		AppBuilder: NewAppBuilderService(config),
	}
}
