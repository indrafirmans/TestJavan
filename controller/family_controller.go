package controller

import (
	"javanid/dto"
	"javanid/model"
	"javanid/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var payloadFamily struct {
	Name     string
	ParentId int `json:"parent"`
}

func AddFamily(c *gin.Context) {
	c.Bind(&payloadFamily)
	status := http.StatusCreated
	message := http.StatusText(status)
	detail := "Data created"
	params := model.Family{
		Name:     payloadFamily.Name,
		ParentId: payloadFamily.ParentId,
	}

	if result := service.AddFamily(&params); result.Error != nil {
		status = http.StatusBadRequest
		message = http.StatusText(status)
		detail = result.Error.Error()
	}

	c.JSON(status, dto.Response{
		Status:  status,
		Message: message,
		Detail:  detail,
		Data:    nil,
	})

}

func GetFamily(c *gin.Context) {
	status := http.StatusOK
	message := http.StatusText(status)
	details := "Data found"
	result := service.GetFamily()

	if result == nil {
		details = "Data empty"
	}

	c.JSON(status, dto.Response{
		Status:  status,
		Message: message,
		Detail:  details,
		Data:    result,
	})
}

func GetFamilyById(c *gin.Context) {
	status := http.StatusOK
	message := http.StatusText(status)
	details := "Data found"
	id := c.Param("id")

	newId, err := strconv.Atoi(id)
	if err != nil {
		status = http.StatusBadRequest
		message = http.StatusText(http.StatusBadRequest)
		details = err.Error()
	}

	result := service.GetFamilyById(newId)
	if result.ID == 0 {
		details = "Data not found"
	}

	c.JSON(status, dto.Response{
		Status:  status,
		Message: message,
		Detail:  details,
		Data:    result,
	})
}

func UpdateFamily(c *gin.Context) {
	status := http.StatusOK
	message := http.StatusText(status)
	details := "Data modified"
	id := c.Param("id")
	c.Bind(&payloadFamily)
	params := model.Family{
		Name:     payloadFamily.Name,
		ParentId: payloadFamily.ParentId,
	}

	newId, err := strconv.Atoi(id)
	if err != nil {
		status = http.StatusBadRequest
		message = http.StatusText(status)
		details = err.Error()
	}

	result := service.UpdateFamily(newId, params)
	if result.Error != nil {
		status = http.StatusBadRequest
		message = http.StatusText(status)
		details = result.Error.Error()
	}

	c.JSON(status, dto.Response{
		Status:  status,
		Message: message,
		Detail:  details,
		Data:    result,
	})
}

func DeleteFamily(c *gin.Context) {
	status := http.StatusOK
	message := http.StatusText(status)
	details := "Data deleted"
	id := c.Param("id")
	c.Bind(&payloadFamily)

	newId, err := strconv.Atoi(id)
	if err != nil {
		status = http.StatusBadRequest
		message = http.StatusText(http.StatusBadRequest)
		details = err.Error()
	}

	result := service.RemoveFamily(newId)
	if result.Error != nil {
		status = http.StatusBadRequest
		message = http.StatusText(status)
		details = result.Error.Error()
	}

	c.JSON(status, dto.Response{
		Status:  status,
		Message: message,
		Detail:  details,
		Data:    result,
	})
}
