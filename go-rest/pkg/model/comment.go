package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserID uint
	Text   string
}
