package router

import (
	"generated_projects/carbon_track/internal/controller"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, ctrl *controller.CarbonController) {
	api := r.Group("/api")
	{
		api.POST("/calculate", ctrl.CalculateFootprint)
		api.GET("/footprints", ctrl.GetFootprintRecords)
	}
}
