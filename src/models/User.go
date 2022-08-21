package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint    `json:"id" gorm:"primarykey"`
	Username  string  `json:"user_name"  valid:"length(4|15),required"`
	Password  []byte  `json:"-" `
	HasAccess bool    `json:"access"`
	Email     string  `json:"user_email"  valid:"email,required"`
	Product   Product `gorm:"-"`
	Comment   Comment `gorm:"-"`
}
