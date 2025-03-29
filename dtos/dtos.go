package dtos

import "time"

type CreateItemRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateItemRequest struct {
	Name string `json:"name"`
}

type SingleItemResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ItemListResponse struct {
	Items []SingleItemResponse `json:"items"`
	// pagination can be added easily here
}

type AddItemToCartRequest struct {
	ItemID       uint    `json:"itemId" binding:"required"`
	Quantity     float32 `json:"quantity" binding:"required"`
	QuantityUnit string  `json:"quantityUnit" binding:"required"`
}

type UpdateCartItemRequest struct {
	Quantity     float32 `json:"quantity"`
	QuantityUnit string  `json:"quantityUnit"`
}

type UpdateCartRequest struct {
	Status string
}