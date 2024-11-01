package routers

import (
	"github.com/gin-gonic/gin"
	"karadyaur.io/ai-dev-light/ad-api-getaway/internal/config"
	"karadyaur.io/ai-dev-light/ad-api-getaway/internal/controller"
)

type appGeneratorRouter struct {
	appGeneratorController *controller.AppGeneratorController
	config                 *config.Config
}

func newAppGeneratorRouter(appGeneratorController *controller.AppGeneratorController, config *config.Config) *appGeneratorRouter {
	return &appGeneratorRouter{appGeneratorController, config}
}

func (agr *appGeneratorRouter) setAppGeneratorRouter(rg *gin.RouterGroup) {
	router := rg.Group("app-generator/")
	router.POST("generate", agr.appGeneratorController.AddRequestForApplicationGenerating)
}
