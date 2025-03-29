package controllers

import (
	"errors"
	"shopping_cart/database"
	"shopping_cart/dtos"
	"shopping_cart/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getActiveCart(createIfNotExist bool) (*models.Cart, error) {
	activeCart := &models.Cart{}
	result := database.DB.Where("status = ?", "active").First(&activeCart)
	if result.RowsAffected <= 0 {
		if !createIfNotExist {
			return nil, errors.New("cart not found")
		}
		activeCart.Status = "active"
		database.DB.Create(&activeCart)
	} else if result.Error != nil {
		return nil, errors.New("unknown error occurred")
	}
	return activeCart, nil
}

func AddItemToCart(c *gin.Context) {
	addItemToCartRequest := dtos.AddItemToCartRequest{}
	if err := c.ShouldBindJSON(&addItemToCartRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	activeCart, err := getActiveCart(true)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	cartItem := models.CartItem{
		CartID:       activeCart.ID,
		ItemID:       addItemToCartRequest.ItemID,
		Quantity:     addItemToCartRequest.Quantity,
		QuantityUnit: addItemToCartRequest.QuantityUnit,
	}

	database.DB.Create(&cartItem)
	c.JSON(201, cartItem)
}

func GetCart(c *gin.Context) {
	activeCart, err := getActiveCart(false)
	if err != nil && err.Error() == "unknown error occurred" {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	} else if err != nil && err.Error() == "cart not found" {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	cartItems := []models.CartItem{}
	database.DB.Where("cart_id = ?", activeCart.ID).Find(&cartItems)
	c.JSON(200, gin.H{"cartItems": cartItems})
}

func DeleteCartItem(c *gin.Context) {
	cartItemId, _ := strconv.Atoi(c.Param("cart_item_id"))
	database.DB.Delete(&models.CartItem{}, cartItemId)

	activeCart, err := getActiveCart(false)
	if err != nil && err.Error() == "unknown error occurred" {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	} else if err != nil && err.Error() == "cart not found" {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	var itemCount int64
	database.DB.Model(models.CartItem{}).Where("cart_id = ?", activeCart.ID).Count(&itemCount)
	if itemCount == 0 {
		updateCartReqeust := dtos.UpdateCartRequest{
			Status: "deleted",
		}
		database.DB.Model(models.Cart{}).Where("id = ?", activeCart.ID).Updates(&updateCartReqeust)
		database.DB.Delete(&activeCart)
	}

	c.JSON(200, gin.H{"message": "Cart item deleted successfully!"})
}

func UpdateCartItem(c *gin.Context) {
	cartItemId, _ := strconv.Atoi(c.Param("cart_item_id"))
	updateCartItemRequest := dtos.UpdateCartItemRequest{}
	if err := c.ShouldBindJSON(&updateCartItemRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	result := database.DB.Model(&models.CartItem{}).Where("id = ?", cartItemId).Updates(&updateCartItemRequest)
	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Cart item not found"})
		return
	}
	c.JSON(200, gin.H{"message": "Cart item updated successfully!"})
}

func CheckoutCart(c *gin.Context) {
	activeCart, err := getActiveCart(false)
	if err != nil && err.Error() == "unknown error occurred" {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	} else if err != nil && err.Error() == "cart not found" {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	updateCartReqeust := &dtos.UpdateCartRequest{
		Status: "order completed",
	}
	result := database.DB.Model(&models.Cart{}).Where("id = ?", activeCart.ID).Updates(&updateCartReqeust)
	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Cart not found"})
		return
	}
	c.JSON(200, gin.H{"message": "Cart checkout done successfully"})
}
