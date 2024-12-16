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
		// CRUD untuk parcel
		parcelRoutes.GET("/", controllers.GetParcels)         // Ambil semua data parcel
		parcelRoutes.GET("/:id", controllers.GetParcelByID)   // Ambil data parcel berdasarkan ID
		parcelRoutes.POST("/", controllers.CreateParcel)      // Tambahkan data parcel baru
		parcelRoutes.PUT("/:id", controllers.UpdateParcel)    // Update data parcel
		parcelRoutes.DELETE("/:id", controllers.DeleteParcel) // Hapus data parcel

		// Routes tambahan untuk tracking status
		parcelRoutes.POST("/:id/tracking", controllers.AddTrackingStatus) // Tambahkan tracking status baru
		parcelRoutes.GET("/:id/tracking", controllers.GetTrackingStatus)  // Ambil riwayat tracking status
	}

	return router
}
