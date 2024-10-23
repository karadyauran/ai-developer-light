package controller

import (
	"generated_projects/carbon_track/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CarbonController struct {
	service *service.CarbonService
}

func NewCarbonController(svc *service.CarbonService) *CarbonController {
	return &CarbonController{service: svc}
}

func (cc *CarbonController) CalculateFootprint(c *gin.Context) {
	var input service.CarbonInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := cc.service.CalculateFootprint(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (cc *CarbonController) GetFootprintRecords(c *gin.Context) {
	records, err := cc.service.GetFootprintRecords()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, records)
}
