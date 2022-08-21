package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID           uint `json:"id" gorm:"primarykey"`
	Category     string  `json:"category"  valid:"required"`
	SubCategory  string  `json:"subcategory"  valid:"required"    `
	Name         string  `json:"name"   valid:"required"   `
	Price        int     `json:"price"  valid:"required"  `
	CommentRefer int     `json:"Comment_id"`
	Comment      Comment `gorm:"foreignKey:CommentRefer"`
}
