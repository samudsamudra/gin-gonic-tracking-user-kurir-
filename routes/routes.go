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
		parcelRoutes.GET("/:id", controllers.GetParcelByID)
		parcelRoutes.POST("/", controllers.CreateParcel)
		parcelRoutes.PUT("/:id", controllers.UpdateParcel)
		parcelRoutes.DELETE("/:id", controllers.DeleteParcel)
		parcelRoutes.POST("/:id/tracking", controllers.AddTrackingStatus)
		parcelRoutes.GET("/:id/tracking", controllers.GetTrackingStatus)
	}

	// Routes untuk user
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/register", controllers.RegisterUser)
		userRoutes.POST("/login", controllers.LoginUser)
	}

	return router
}
