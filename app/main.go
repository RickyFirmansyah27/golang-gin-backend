package main

import (
	"golang-vercel/app/config"
	"golang-vercel/app/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	err := config.DBConnection()
	if err != nil {
		log.Fatal("Database connection failed: ", err)
	}

	router := gin.Default()
	routes.RootRoute(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("[Gin-Service] Server is running on port %s...", port)
	log.Fatal(router.Run(":" + port))
}
