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

// Authenticate handles GitHub OAuth authentication
// @Summary      Authenticate user through GitHub OAuth
// @Description  This endpoint accepts a GitHub authorization code and returns the authenticated user details.
// @Tags         OAuth
// @Accept       json
// @Produce      json
// @Param        request  body      model.AuthenticateUserRequestSwagger  true  "GitHub Authorization Code"
// @Success      200      {object}  model.UserResponseSwagger  "User response with avatar"
// @Failure      400      {object}  map[string]string  "Invalid request"
// @Failure      500      {object}  map[string]string  "Internal server error"
// @Router       /api/v1/oauth/github/authenticate [post]
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
