package service

import (
	"javanid/config/db"
	"javanid/model"

	"gorm.io/gorm"
)

func AddFamily(params *model.Family) *gorm.DB {
	return db.Conn().Create(&params)
}

func GetFamily() []model.Family {
	var members []model.Family

	db.Conn().Find(&members)

	return members
}

func GetFamilyById(id int) model.Family {
	var members model.Family

	db.Conn().First(&members, id)

	return members
}

func UpdateFamily(id int, params model.Family) *gorm.DB {
	var family model.Family
	return db.Conn().Model(family).Where("id = ?", id).Updates(model.Family{
		Name:     params.Name,
		ParentId: params.ParentId,
	})
}

func RemoveFamily(id int) *gorm.DB {
	return db.Conn().Delete(&model.Family{}, id)
}
