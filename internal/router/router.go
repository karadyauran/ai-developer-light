package routers

import (
	"ai-dev-light/internal/config"
	"ai-dev-light/internal/controller"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

type Router struct {
	Gin    *gin.Engine
	config *config.Config
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
		Gin:    ginRouter,
		config: config,
	}
}

func (r *Router) SetRoutes() {
	if r.config.EnvType != "prod" {
		r.Gin.GET("/healthcheck", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "The way APi is working fine"})
		})
		r.Gin.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
