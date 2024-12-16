package routes

import (
	"sistem-tracking/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	// Routes untuk parcel
	parcelRoutes := router.Group("/parcels")
	{
		parcelRoutes.GET("/", controllers.GetParcels)
		parcelRoutes.POST("/", controllers.CreateParcel)
		parcelRoutes.PUT("/:id", controllers.UpdateParcel)
		parcelRoutes.DELETE("/:id", controllers.DeleteParcel)
		parcelRoutes.GET("/:id", controllers.GetParcelByID)
	}

	return router
}
