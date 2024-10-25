package routers

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"karadyaur.io/ai-dev-light/ad-api-getaway/internal/config"
	"karadyaur.io/ai-dev-light/ad-api-getaway/internal/controller"
	"net/http"
)

type Router struct {
	Gin         *gin.Engine
	config      *config.Config
	oAuthRouter *oAuthRouter
}

func NewRouter(config *config.Config, controller *controller.Controller) *Router {
	ginRouter := gin.Default()

	// Apply CORS middleware with custom options
	ginRouter.Use(cors.New(cors.Config{
		AllowOrigins:     []string{config.WebappBaseUrl},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	ginRouter.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": fmt.Sprintf("The specified route %s not found", ctx.Request.URL)})
	})

	return &Router{
		Gin:         ginRouter,
		config:      config,
		oAuthRouter: newOAuthRouter(controller.OAuthController, config),
	}
}

func (r *Router) SetRoutes() {
	api := r.Gin.Group("/api/v1")

	r.oAuthRouter.setOAuthRoutes(api)

	if r.config.EnvType != "prod" {
		// r.devRouter.setDevRoutes(api)
		r.Gin.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
