package service

import (
	"javanid/config/db"
	"javanid/model"
)

func GetMembers() []model.Family {
	var members []model.Family
	db.Conn().Find(&members)

	return members
}
