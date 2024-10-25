package controller

import "karadyaur.io/ai-dev-light/ad-api-getaway/internal/generated"

type Controller struct {
	OAuthController *OAuthController
}

func NewController(authClient generated.OAuthServiceClient) *Controller {
	return &Controller{
		OAuthController: NewAuthController(authClient),
	}
}
