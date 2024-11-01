package controller

import "karadyaur.io/ai-dev-light/ad-api-getaway/internal/generated"

type Controller struct {
	OAuthController        *OAuthController
	AppGeneratorController *AppGeneratorController
}

func NewController(authClient generated.OAuthServiceClient, kafkaClient generated.KafkaServiceClient) *Controller {
	return &Controller{
		OAuthController:        NewAuthController(authClient),
		AppGeneratorController: NewAppGeneratorController(kafkaClient),
	}
}
