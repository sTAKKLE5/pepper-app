package routes

import (
	"github.com/gin-gonic/gin"
	"os"
	"pepper-app/internal/handlers"
)

func SetupRouter() *gin.Engine {
	mode := os.Getenv("MODE")

	switch mode {
	case "release":
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	router := gin.Default()

	plantsGroup := router.Group("/plants")
	{
		// Plants routes
		plantsGroup.GET("", handlers.GetPlants)
		plantsGroup.GET("/:id", handlers.GetPlantByID)

		plantsGroup.POST("", handlers.CreatePlant)
		plantsGroup.PUT("/:id", handlers.UpdatePlantByID)
		plantsGroup.DELETE("/:id", handlers.DeletePlantByID)

		return router
	}
}
