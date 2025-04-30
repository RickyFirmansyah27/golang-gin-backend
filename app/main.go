package main

import (
	"golang-vercel/app/config"
	"golang-vercel/app/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Database connection
	err := config.DBConnection()
	if err != nil {
		log.Fatal("Database connection failed: ", err)
	}

	// Create Gin router
	router := gin.Default()

	// Setup routes
	routes.RootRoute(router)

	// Get port from environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Start server
	log.Printf("Gin server started on %s...", port)
	log.Fatal(router.Run(":" + port))
}
