package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Description   string   `json:"description"  valid:"required"`
	ReplyID       *uint
	Reply	     []Comment `valid:"-"  gorm:"foreignkey:ReplyID"`
	UserID         uint
}
