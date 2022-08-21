package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	ID          uint   `json:"id" gorm:"primarykey"`
	Description string `json:"desc"`
}
