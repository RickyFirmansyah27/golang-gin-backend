package service

import (
	"golang-vercel/app/config"
	"golang-vercel/app/models"
	"log"

	"github.com/gin-gonic/gin"
)

func GetItems(c *gin.Context) (int, []models.Item, error) {
	log.Println("[ItemsService] - Fetching items...", c.Request.URL.Query())

	// Get query parameters from Gin context
	queryParams := make(map[string]string)
	for key, values := range c.Request.URL.Query() {
		if len(values) > 0 {
			queryParams[key] = values[0]
		}
	}

	itemsMap, totalData, err := config.GetAllItems(queryParams)
	if err != nil {
		log.Printf("[ItemsService] - Error fetching items: %v", err)
		return 0, nil, err
	}

	items := make([]models.Item, 0, len(itemsMap))
	for _, itemMap := range itemsMap {
		item := models.Item{
			ID:         itemMap["id"].(int),
			Name:       itemMap["name"].(string),
			CategoryID: itemMap["category_id"].(int),
			Stock:      itemMap["stock"].(int),
			Unit:       itemMap["unit"].(string),
			MinStock:   itemMap["min_stock"].(int),
		}
		items = append(items, item)
	}

	log.Printf("[ItemsService] - Successfully fetched %d items", len(items))
	return totalData, items, nil
}
