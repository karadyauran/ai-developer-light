package routers

import (
	"github.com/gin-gonic/gin"
	"karadyaur.io/ai-dev-light/ad-api-getaway/internal/config"
	"karadyaur.io/ai-dev-light/ad-api-getaway/internal/controller"
)

type oAuthRouter struct {
	oAuthController *controller.OAuthController
	config          *config.Config
}

func newOAuthRouter(authController *controller.OAuthController, config *config.Config) *oAuthRouter {
	return &oAuthRouter{authController, config}
}

func (ar *oAuthRouter) setOAuthRoutes(rg *gin.RouterGroup) {
	router := rg.Group("oauth/github")
	router.POST("/authenticate", ar.oAuthController.Authenticate)
}
