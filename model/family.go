package model

import "gorm.io/gorm"

type Family struct {
	gorm.Model
	Name     string
	ParentId uint
}
