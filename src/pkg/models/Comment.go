package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Description string `json:"desc"`
}
