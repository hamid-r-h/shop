package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string  `json:"username"  valid:"length(4|15),required"`
	Password  string  `json:"password"  valid:"length(4|15),required" `
	HasAccess bool    `json:"access"`
	Email     string  `json:"email"  valid:"email,required"`
	Product   Product `valid:"-"`
}
