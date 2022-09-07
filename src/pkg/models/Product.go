package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Category    string ` valid:"required"`
	SubCategory string ` valid:"required"`
	Name        string ` valid:"required" gorm:"unique"   `
	Price       int    ` valid:"required"`
	Comments  []Comment  `valid:"-"  gorm:"many2many:product_comments;"`
	Number      int    ` valid:"required"`
}