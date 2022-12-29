package controller

import (
	"javanid/config/db"
	"javanid/model"
	"javanid/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAsset(c *gin.Context) {

	members := service.GetMembers()

	if members == nil {
		Status = http.StatusBadRequest
	}

	c.JSON(Status, gin.H{
		"status": Status,
		"data":   members,
	})
}

func AddAsset(c *gin.Context) {
	c.Bind(&payload)
	message = http.StatusText(http.StatusCreated)
	Status = http.StatusCreated
	params := model.Family{
		Name: payload.Name,
	}

	result := db.Conn().Create(&params)

	if result.Error != nil {
		Status = http.StatusBadRequest
		message = result.Error.Error()
	}

	c.JSON(Status, gin.H{
		"message": message,
		"data":    params,
	})
}

func UpdateAsset(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello Boss",
	})
}

func DeleteAsset(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello Boss",
	})
}
