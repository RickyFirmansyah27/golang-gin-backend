package controllers

import (
	"log"
	"net/http"

	"golang-vercel/app/helpers"
	"golang-vercel/app/models"
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

func CreateItem(c *gin.Context) {
	var newItem models.Item
	if err := c.ShouldBindJSON(&newItem); err != nil {
		log.Printf("[ItemsController] - Invalid input: %v", err)
		helpers.Error(c, http.StatusBadRequest, "Invalid input", err)
		return
	}

	err := service.CreateItem(newItem)
	if err != nil {
		log.Printf("[ItemsController] - Failed to create item: %v", err)
		helpers.Error(c, http.StatusInternalServerError, "Failed to create item", err)
		return
	}

	log.Printf("[ItemsController] - Successfully created item")
	helpers.Success(c, "Successfully created item", gin.H{"item": newItem})
}

func UpdateItem(c *gin.Context) {
	var updatedItem models.Item
	if err := c.ShouldBindJSON(&updatedItem); err != nil {
		log.Printf("[ItemsController] - Invalid input: %v", err)
		helpers.Error(c, http.StatusBadRequest, "Invalid input", err)
		return
	}

	err := service.UpdateItem(updatedItem)
	if err != nil {
		log.Printf("[ItemsController] - Failed to update item: %v", err)
		helpers.Error(c, http.StatusInternalServerError, "Failed to update item", err)
		return
	}

	log.Printf("[ItemsController] - Successfully updated item")
	helpers.Success(c, "Successfully updated item", gin.H{"item": updatedItem})
}

func DeleteItem(c *gin.Context) {
	id := c.Param("id")

	err := service.DeleteItem(id)
	if err != nil {
		log.Printf("[ItemsController] - Failed to delete item: %v", err)
		helpers.Error(c, http.StatusInternalServerError, "Failed to delete item", err)
		return
	}

	log.Printf("[ItemsController] - Successfully deleted item")
	helpers.Success(c, "Successfully deleted item", nil)
}
