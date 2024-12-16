package main

import (
	"fmt"
	"log"
	"net/http"
	"sistem-tracking/config"
	"sistem-tracking/routes"
)

func main() {
	// Hubungkan ke database
	config.ConnectDB()

	// Definisikan route
	router := routes.SetupRoutes()

	// Jalankan server
	port := ":8080"
	fmt.Printf("cihuyyy, servernya jalan coo, jalan di http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
