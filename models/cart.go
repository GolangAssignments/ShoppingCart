package models 
import (
	"time"
	"gorm.io/gorm"
)

type Cart struct {
	ID uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Status string `gorm:"size:255"`
}