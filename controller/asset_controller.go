package controller

import (
	"javanid/dto"
	"javanid/model"
	"javanid/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var payloadAsset struct {
	Name       string
	AssetOwner int `json:"owner"`
}

func AddAsset(c *gin.Context) {
	c.Bind(&payloadAsset)
	status := http.StatusCreated
	message := http.StatusText(status)
	detail := "Data created"
	params := model.Asset{
		Name:       payloadAsset.Name,
		AssetOwner: payloadAsset.AssetOwner,
	}

	result := service.AddAsset(&params)

	if result.Error != nil {
		status = http.StatusBadRequest
		message = result.Error.Error()
	}

	c.JSON(status, dto.Response{
		Status:  status,
		Message: message,
		Detail:  detail,
		Data:    nil,
	})
}

func GetAsset(c *gin.Context) {
	status := http.StatusOK
	message := http.StatusText(status)
	details := "Data found"
	result := service.GetAsset()

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

func GetAssetById(c *gin.Context) {
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

	result := service.GetAssetById(newId)
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

func UpdateAsset(c *gin.Context) {
	status := http.StatusOK
	message := http.StatusText(status)
	details := "Data modified"
	id := c.Param("id")
	c.Bind(&payloadAsset)
	params := model.Asset{
		Name:       payloadFamily.Name,
		AssetOwner: payloadAsset.AssetOwner,
	}

	newId, err := strconv.Atoi(id)

	if err != nil {
		status = http.StatusBadRequest
		message = http.StatusText(http.StatusBadRequest)
		details = err.Error()
	}

	result := service.UpdateAsset(newId, params)
	if result.Error != nil {
		status = http.StatusBadRequest
		message = result.Error.Error()
		details = result.Error.Error()
	}

	c.JSON(status, dto.Response{
		Status:  status,
		Message: message,
		Detail:  details,
		Data:    nil,
	})
}

func DeleteAsset(c *gin.Context) {
	status := http.StatusOK
	message := http.StatusText(status)
	details := "Data deleted"
	id := c.Param("id")
	c.Bind(&payloadAsset)

	newId, err := strconv.Atoi(id)

	if err != nil {
		status = http.StatusBadRequest
		message = http.StatusText(http.StatusBadRequest)
	}

	result := service.RemoveAsset(newId)
	if result.Error != nil {
		status = http.StatusBadRequest
		message = result.Error.Error()
		details = err.Error()
	}

	c.JSON(status, dto.Response{
		Status:  status,
		Message: message,
		Detail:  details,
		Data:    nil,
	})
}
