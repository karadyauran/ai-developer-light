package service

import (
	"ai-dev-light/internal/config"
	"fmt"
)

type Service struct {
	AppBuilder *AppBuilder
}

func NewService(cfg *config.Config) (*Service, error) {
	appBuilder, err := NewAppBuilderService(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create AppBuilder: %w", err)
	}

	return &Service{
		AppBuilder: appBuilder,
	}, nil
}
