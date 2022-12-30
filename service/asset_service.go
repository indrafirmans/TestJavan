package service

import (
	"javanid/config/db"
	"javanid/model"

	"gorm.io/gorm"
)

func AddAsset(params *model.Asset) *gorm.DB {
	return db.Conn().Create(&params)
}

func GetAsset() []model.Asset {
	var data []model.Asset

	db.Conn().Find(&data)

	return data
}

func GetAssetById(id int) model.Asset {
	var datas model.Asset

	db.Conn().First(&datas, id)

	return datas
}

func UpdateAsset(id int, params model.Asset) *gorm.DB {
	var datas model.Asset
	return db.Conn().Model(datas).Where("id = ?", id).Updates(model.Asset{
		Name:       params.Name,
		AssetOwner: params.AssetOwner,
	})
}

func RemoveAsset(id int) *gorm.DB {
	return db.Conn().Delete(&model.Asset{}, id)
}
