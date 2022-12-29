package main

import (
	"javanid/config"
	"javanid/config/db"
	"javanid/model"
)

func init() {
	config.LoadEnvirontment()
	db.Conn()
}

func main() {
	db.Conn().AutoMigrate(
		&model.Family{},
		&model.Asset{},
	)
}
