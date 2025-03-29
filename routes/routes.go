package routes

import (
	"shopping_cart/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	r.POST("/items", controllers.CreateItem)
	r.GET("/items", controllers.GetItems)
	r.PATCH("/items/:id", controllers.UpdateItem)
	r.DELETE("/items/:id", controllers.DeleteItem)
	r.POST("/cart/items", controllers.AddItemToCart)
	r.GET("/cart", controllers.GetCart)
	r.DELETE("/cart/items/:cart_item_id", controllers.DeleteCartItem)
	r.PATCH("/cart/items/:cart_item_id", controllers.UpdateCartItem)
	r.POST("/cart/checkout", controllers.CheckoutCart)
	return r
}
