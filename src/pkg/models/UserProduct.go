package models

import "gorm.io/gorm"

type UserProduct struct {
	gorm.Model
	UserID    int `gorm:"primaryKey"`
	ProductID int `gorm:"primaryKey"`
	Number    int
}
