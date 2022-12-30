package controller

import (
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
	params := model.Asset{
		Name:       payloadAsset.Name,
		AssetOwner: payloadAsset.AssetOwner,
	}

	result := service.AddAsset(&params)

	if result.Error != nil {
		status = http.StatusBadRequest
		message = result.Error.Error()
	}

	c.JSON(status, gin.H{
		"status":  status,
		"message": message,
		"data":    nil,
	})
}

func GetAsset(c *gin.Context) {
	status := http.StatusOK
	message := http.StatusText(status)

	datas := service.GetAsset()

	c.JSON(status, gin.H{
		"status":  status,
		"message": message,
		"data":    datas,
	})
}

func GetAssetById(c *gin.Context) {
	status := http.StatusOK
	message := http.StatusText(status)
	var data model.Asset
	id := c.Param("id")
	newId, err := strconv.Atoi(id)

	if err != nil {
		status = http.StatusBadRequest
		message = http.StatusText(http.StatusBadRequest)
	}

	data = service.GetAssetById(newId)

	c.JSON(status, gin.H{
		"status":  status,
		"message": message,
		"data":    data,
	})
}

func UpdateAsset(c *gin.Context) {
	status := http.StatusOK
	message := http.StatusText(status)
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
	}

	result := service.UpdateAsset(newId, params)
	if result.Error != nil {
		status = http.StatusBadRequest
		message = result.Error.Error()
	}

	c.JSON(status, gin.H{
		"status":  status,
		"message": message,
		"data":    nil,
	})
}

func DeleteAsset(c *gin.Context) {
	status := http.StatusOK
	message := http.StatusText(status)
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
	}

	c.JSON(status, gin.H{
		"status":  status,
		"message": message,
		"data":    nil,
	})
}
