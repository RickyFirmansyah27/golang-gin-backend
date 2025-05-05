package service

import (
	"golang-vercel/app/config"
	"golang-vercel/app/models"
	"log"

	"github.com/gin-gonic/gin"
)

func GetItems(c *gin.Context) (int, []models.Item, error) {
	log.Println("[ItemsService] - Fetching items...", c.Request.URL.Query())

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

func CreateItem(newItem models.Item) error {
	log.Println("[ItemsService] - Creating item...")

	err := config.CreateItem(newItem)
	if err != nil {
		log.Printf("[ItemsService] - Error creating item: %v", err)
		return err
	}

	log.Printf("[ItemsService] - Successfully created item: %v", newItem)
	return nil
}

func UpdateItem(updatedItem models.Item) error {
	log.Println("[ItemsService] - Updating item...")
	log.Println(updatedItem)

	err := config.UpdateItem(updatedItem)
	if err != nil {
		log.Printf("[ItemsService] - Error updating item: %v", err)
		return err
	}

	log.Printf("[ItemsService] - Successfully updated item: %v", updatedItem)
	return nil
}

func DeleteItem(id string) error {
	log.Println("[ItemsService] - Deleting item...")

	err := config.DeleteItem(id)
	if err != nil {
		log.Printf("[ItemsService] - Error deleting item: %v", err)
		return err
	}

	log.Printf("[ItemsService] - Successfully deleted item: %v", id)
	return nil
}
