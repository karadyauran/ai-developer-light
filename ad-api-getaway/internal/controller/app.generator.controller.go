package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"karadyaur.io/ai-dev-light/ad-api-getaway/internal/generated"
	"net/http"
	"time"
)

type AppGeneratorController struct {
	client generated.KafkaServiceClient
}

func NewAppGeneratorController(kafkaClient generated.KafkaServiceClient) *AppGeneratorController {
	return &AppGeneratorController{
		client: kafkaClient,
	}
}
func (c *AppGeneratorController) AddRequestForApplicationGenerating(gc *gin.Context) {
	var req generated.KafkaRequest

	if err := gc.ShouldBindJSON(&req); err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if req.Topic == "" {
		gc.JSON(http.StatusBadRequest, gin.H{"error": "Topic is empty"})
	}

	if req.Params.GenerationType == "" {
		gc.JSON(http.StatusBadRequest, gin.H{"error": "Parameter generationType is empty"})
	}

	if req.Params.Token == "" {
		gc.JSON(http.StatusBadRequest, gin.H{"error": "Parameter token is empty"})
	}

	ctx := context.Background()

	req.Id = "213"
	req.CreatedAt = time.Now().String()

	response, err := c.client.KafkaSend(ctx, &req)
	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, response)
}
