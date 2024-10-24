package router

import (
	"eco_track_api/internal/config"
	"eco_track_api/internal/controller"
	"github.com/gin-gonic/gin"
)

type Router struct {
	config    *config.Config
	ginEngine *gin.Engine
	controller *controller.CarbonController
}

func NewRouter(cfg *config.Config, carbonCtrl *controller.CarbonController) *Router {
	return &Router{
		config:    cfg,
		ginEngine: gin.Default(),
		controller: carbonCtrl,
	}
}

func (r *Router) SetRoutes() {
	api := r.ginEngine.Group("/api/v1")
	{
		api.POST("/carbon/calculate", r.controller.CalculateCarbon)
	}
}

func (r *Router) Gin() *gin.Engine {
	return r.ginEngine
}