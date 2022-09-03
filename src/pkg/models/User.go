package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string    `json:"username"  valid:"length(4|15),required"  gorm:"unique"`
	Password  string    `json:"password"  valid:"length(4|15),required" `
	HasAccess bool      `json:"access"  `
	Email     string    `json:"email"  valid:"email,required"`
	
	Products  []Product `valid:"-" gorm:"many2many:user_products;"`
}
