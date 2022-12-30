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
	var datas []model.Family

	db.Conn().Find(&datas)

	return datas
}

func GetFamilyById(id int) model.Family {
	var datas model.Family

	db.Conn().First(&datas, id)

	return datas
}

func UpdateFamily(id int, params model.Family) *gorm.DB {
	var datas model.Family
	return db.Conn().Model(datas).Where("id = ?", id).Updates(model.Family{
		Name:     params.Name,
		ParentId: params.ParentId,
	})
}

func RemoveFamily(id int) *gorm.DB {
	return db.Conn().Delete(&model.Family{}, id)
}
