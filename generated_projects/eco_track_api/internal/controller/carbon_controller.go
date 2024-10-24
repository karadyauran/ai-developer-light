package controller

import (
	"net/http"
	"eco_track_api/internal/service"
	"github.com/gin-gonic/gin"
)

type CarbonController struct {
	service *service.CarbonService
}

func NewCarbonController(s *service.CarbonService) *CarbonController {
	return &CarbonController{service: s}
}

func (cc *CarbonController) CalculateCarbon(c *gin.Context) {
	var request service.CarbonCalculationRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := cc.service.Calculate(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}