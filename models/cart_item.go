package models 
import (
	"time"
	"gorm.io/gorm"
)

type CartItem struct {
	ID uint `gorm:"primarykey"`
	CartID uint
	ItemID uint
	Quantity float32
	QuantityUnit string `gorm:"size:20"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}