package controller

import (
	"ai-dev-light/internal/service"
)

type Controller struct {
}

func NewController(service *service.Service) *Controller {
	return &Controller{}
}
