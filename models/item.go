package models 

import (
	"time"
	"gorm.io/gorm"
)

type Item struct {
	ID uint `gorm:"primarykey"`
	Name string `gorm:"size:255"`
	CreatedAt time.Time 
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}