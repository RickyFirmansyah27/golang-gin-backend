package routes

import (
	"golang-vercel/app/controllers"
	"golang-vercel/app/helpers"

	"net/http"

	"github.com/gin-gonic/gin"
)

func RootRoute(app *gin.Engine) {
	app.NoRoute(noAccessRoute)
	app.GET("/items", controllers.GetAllItems)

}

func noAccessRoute(c *gin.Context) {
	c.JSON(http.StatusOK, helpers.BaseResponse{
		Success: true,
		Message: "No Route Found",
	})
}
