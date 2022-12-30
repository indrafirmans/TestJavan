package controller

import (
	"javanid/model"
	"javanid/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Status int = http.StatusOK
var message string

var payload struct {
	Name     string
	ParentId int `json:"parent"`
}

func AddFamily(c *gin.Context) {
	c.Bind(&payload)
	Status = http.StatusCreated
	message = http.StatusText(Status)
	params := model.Family{
		Name:     payload.Name,
		ParentId: uint(payload.ParentId),
	}

	result := service.AddFamily(&params)

	if result.Error != nil {
		Status = http.StatusBadRequest
		message = result.Error.Error()
	}

	c.JSON(Status, gin.H{
		"status":  Status,
		"message": message,
		"data":    nil,
	})
}

func GetFamily(c *gin.Context) {

	members := service.GetFamily()

	c.JSON(Status, gin.H{
		"status":  Status,
		"message": http.StatusText(Status),
		"data":    members,
	})
}

func GetFamilyById(c *gin.Context) {
	var members model.Family
	id := c.Param("id")
	message := http.StatusText(Status)
	newId, err := strconv.Atoi(id)

	if err != nil {
		Status = http.StatusBadRequest
		message = http.StatusText(http.StatusBadRequest)
	}

	members = service.GetFamilyById(newId)

	c.JSON(Status, gin.H{
		"status":  Status,
		"message": message,
		"data":    members,
	})
}

func UpdateFamily(c *gin.Context) {
	id := c.Param("id")
	c.Bind(&payload)
	message = "record modified"
	params := model.Family{
		Name:     payload.Name,
		ParentId: uint(payload.ParentId),
	}

	newId, err := strconv.Atoi(id)

	if err != nil {
		Status = http.StatusBadRequest
		message = http.StatusText(http.StatusBadRequest)
	}

	result := service.UpdateFamily(newId, params)
	if result.Error != nil {
		Status = http.StatusBadRequest
		message = result.Error.Error()
	}

	c.JSON(Status, gin.H{
		"status":  Status,
		"message": message,
		"data":    nil,
	})
}

func DeleteFamily(c *gin.Context) {
	id := c.Param("id")
	c.Bind(&payload)
	message = "record deleted"

	newId, err := strconv.Atoi(id)

	if err != nil {
		Status = http.StatusBadRequest
		message = http.StatusText(http.StatusBadRequest)
	}

	result := service.RemoveFamily(newId)
	if result.Error != nil {
		Status = http.StatusBadRequest
		message = result.Error.Error()
	}

	c.JSON(Status, gin.H{
		"status":  Status,
		"message": message,
		"data":    nil,
	})
}
