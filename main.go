package main

import (
	"fmt"
	"log"
	"sistem-tracking/config"
	"sistem-tracking/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Mode Release untuk production
	gin.SetMode(gin.ReleaseMode)

	// Hubungkan ke database
	config.ConnectDB()

	// Log untuk memastikan migrasi berjalan
	fmt.Println("Migrasi database sedang dijalankan...")

	// Setup router
	router := routes.SetupRoutes()
	router.SetTrustedProxies(nil)

	// Port server
	port := ":8080"
	fmt.Printf("Server berjalan di http://localhost%s\n", port)

	// Jalankan server
	if err := router.Run(port); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
