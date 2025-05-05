package handler

import (
	"log"
	"net/http"

	"golang-vercel/app/config"
	"golang-vercel/app/routes"

	"github.com/gin-contrib/cors"
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

	// Add CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	routes.RootRoute(app)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
