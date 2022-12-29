package controller

import (
	"javanid/config/db"
	"javanid/model"
	"javanid/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Status int = http.StatusOK
var message string

var payload struct {
	Name string
}

func GetFamily(c *gin.Context) {

	members := service.GetMembers()

	if members == nil {
		Status = http.StatusBadRequest
	}

	c.JSON(Status, gin.H{
		"status": Status,
		"data":   members,
	})
}

func AddFamily(c *gin.Context) {
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

func UpdateFamily(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello Boss",
	})
}

func DeleteFamily(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello Boss",
	})
}
