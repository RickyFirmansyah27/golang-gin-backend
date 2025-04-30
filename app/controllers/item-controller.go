package controllers

import (
	"log"
	"net/http"

	"golang-vercel/app/helpers"
	"golang-vercel/app/service"

	"github.com/gin-gonic/gin"
)

func GetAllItems(c *gin.Context) {
	log.Printf("[ItemsController] - Incoming request with query params: %v", c.Request.URL.Query())

	totalData, items, err := service.GetItems(c)
	if err != nil {
		log.Printf("[ItemsController] - Failed to fetch items: %v", err)
		helpers.Error(c, http.StatusBadRequest, "Failed to fetch items", err)
		return
	}

	response := gin.H{
		"total_data": totalData,
		"items":      items,
	}

	log.Printf("[ItemsController] - Successfully fetched items")
	helpers.Success(c, "Successfully fetched items", response)
}
