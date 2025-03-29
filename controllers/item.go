package controllers

import (
	"shopping_cart/database"
	"shopping_cart/dtos"
	"shopping_cart/models"
	strconv "strconv"

	"github.com/gin-gonic/gin"
)

func CreateItem(c *gin.Context) {
	var createItemRequest dtos.CreateItemRequest
	if err := c.ShouldBindJSON(&createItemRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	item := models.Item{
		Name: createItemRequest.Name,
	}

	database.DB.Create(&item)
	c.JSON(201, dtos.SingleItemResponse{
		ID:        item.ID,
		Name:      item.Name,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	})
}

func GetItems(c *gin.Context) {
	itemList := []models.Item{}
	database.DB.Find(&itemList)

	itemResponseList := dtos.ItemListResponse{}

	for _, item := range itemList {
		singleItemResponse := dtos.SingleItemResponse{
			ID:        item.ID,
			Name:      item.Name,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		}
		itemResponseList.Items = append(itemResponseList.Items, singleItemResponse)
	}
	c.JSON(200, itemResponseList)
}

func UpdateItem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	updateItemRequest := dtos.UpdateItemRequest{}
	if err := c.ShouldBindJSON(&updateItemRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	database.DB.Model(&models.Item{}).Where("id = ?", id).Updates(&updateItemRequest)
	c.JSON(200, gin.H{"message": "Item updated successfully!"})
}

func DeleteItem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	tx := database.DB.Delete(&models.Item{}, id)
	if tx.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Item not found!"})
		return
	}
	c.JSON(200, gin.H{"message": "Item deleted successfully!"})
}
