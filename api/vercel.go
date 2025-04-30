package handler

import (
	"log"
	"net/http"

	"golang-vercel/app/config"
	"golang-vercel/app/routes"

	"github.com/gin-gonic/gin"
)

var (
	app *gin.Engine
)

func init() {
	gin.SetMode(gin.ReleaseMode)

	// Initialize app with middleware
	app = gin.New()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())

	// Setup database connection
	err := config.DBConnection()
	if err != nil {
		log.Printf("Database connection failed: %v", err)
	}

	routes.RootRoute(app)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
