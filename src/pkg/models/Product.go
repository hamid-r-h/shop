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
	Number      int	   ` valid:"required"`
	UserID      uint
}
