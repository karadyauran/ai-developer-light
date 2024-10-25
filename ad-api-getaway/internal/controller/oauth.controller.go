package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"karadyaur.io/ai-dev-light/ad-api-getaway/internal/generated"
	"net/http"
)

type OAuthController struct {
	client generated.OAuthServiceClient
}

func NewAuthController(oAuthClient generated.OAuthServiceClient) *OAuthController {
	return &OAuthController{
		client: oAuthClient,
	}
}

func (c *OAuthController) Authenticate(gc *gin.Context) {
	var req generated.AuthenticateUserRequest

	if err := gc.ShouldBindJSON(&req); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if req.Code == "" {
		gc.JSON(http.StatusBadRequest, gin.H{"error": "code is required"})
		return
	}

	ctx := context.Background()

	response, err := c.client.Authenticate(ctx, &req)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, response)
}
