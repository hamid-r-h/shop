package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Category    string `json:"category"  valid:"required"`
	SubCategory string `json:"subcategory"  valid:"required"`
	Name        string `json:"name"   valid:"required"`
	Price       int    `json:"price"  valid:"required"`
	UserID uint
}
