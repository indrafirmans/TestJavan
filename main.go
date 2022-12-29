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
	r.GET(common.UrlApi+common.UrlFamily, controller.GetFamily)
	r.POST(common.UrlApi+common.UrlFamily, controller.AddFamily)
	r.PUT(common.UrlApi+common.UrlFamily, controller.UpdateFamily)
	r.DELETE(common.UrlApi+common.UrlFamily, controller.DeleteFamily)

	// api asset
	r.GET(common.UrlApi+common.UrlAsset, controller.GetAsset)
	r.POST(common.UrlApi+common.UrlAsset, controller.AddAsset)
	r.PUT(common.UrlApi+common.UrlAsset, controller.UpdateAsset)
	r.DELETE(common.UrlApi+common.UrlAsset, controller.DeleteAsset)

	r.Run(port)
}
