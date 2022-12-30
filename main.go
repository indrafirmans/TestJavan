package main

import (
	"javanid/common"
	"javanid/config"
	"javanid/config/db"
	"javanid/controller"

	"os"

	"github.com/gin-gonic/gin"
)

type Result struct {
	ID   int
	Name string
	Age  int
}

func init() {
	config.LoadEnvirontment()
	db.Conn()
}

func main() {
	port := ":" + os.Getenv("APP_PORT")
	r := gin.Default()

	// api family
	r.POST(common.URL_API+common.URL_FAMILY, controller.AddFamily)
	r.GET(common.URL_API+common.URL_FAMILY, controller.GetFamily)
	r.GET(common.URL_API+common.URL_FAMILY+"/:id", controller.GetFamilyById)
	r.PUT(common.URL_API+common.URL_FAMILY+"/:id", controller.UpdateFamily)
	r.DELETE(common.URL_API+common.URL_FAMILY+"/:id", controller.DeleteFamily)

	// api asset
	r.POST(common.URL_API+common.URL_ASSET, controller.AddAsset)
	r.GET(common.URL_API+common.URL_ASSET, controller.GetAsset)
	r.GET(common.URL_API+common.URL_ASSET+"/:id", controller.GetAssetById)
	r.PUT(common.URL_API+common.URL_ASSET+"/:id", controller.UpdateAsset)
	r.DELETE(common.URL_API+common.URL_ASSET+"/:id", controller.DeleteAsset)

	r.Run(port)
}
