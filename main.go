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
	r.POST(common.UrlApi+common.UrlFamily, controller.AddFamily)
	r.GET(common.UrlApi+common.UrlFamily, controller.GetFamily)
	r.GET(common.UrlApi+common.UrlFamily+"/:id", controller.GetFamilyById)
	r.PUT(common.UrlApi+common.UrlFamily+"/:id", controller.UpdateFamily)
	r.DELETE(common.UrlApi+common.UrlFamily+"/:id", controller.DeleteFamily)

	// api asset
	r.POST(common.UrlApi+common.UrlAsset, controller.AddAsset)
	r.GET(common.UrlApi+common.UrlAsset, controller.GetAsset)
	r.GET(common.UrlApi+common.UrlAsset+"/:id", controller.GetAssetById)
	r.PUT(common.UrlApi+common.UrlAsset+"/:id", controller.UpdateAsset)
	r.DELETE(common.UrlApi+common.UrlAsset+"/:id", controller.DeleteAsset)

	r.Run(port)
}
