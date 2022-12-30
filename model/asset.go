package model

import "gorm.io/gorm"

type Asset struct {
	gorm.Model
	Name       string
	AssetOwner int
}
