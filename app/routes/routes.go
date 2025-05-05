package routes

import (
	"golang-vercel/app/controllers"
	"golang-vercel/app/helpers"
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
)

func RootRoute(app *gin.Engine) {
	app.NoRoute(noAccessRoute)

	// Add logging here for debugging
	log.Println("Defining routes...")
	app.GET("/items", controllers.GetAllItems)
	app.POST("/items", controllers.CreateItem)
	app.PATCH("/items", controllers.UpdateItem)
	app.DELETE("/items", controllers.DeleteItem)
}

func noAccessRoute(c *gin.Context) {
	c.JSON(http.StatusOK, helpers.BaseResponse{
		Success: true,
		Message: "No Route Found",
	})
}
