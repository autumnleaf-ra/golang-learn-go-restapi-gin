package main

import (
	"github.com/autumnleaf-ra/golang/go-restapi-gin/controllers/productcontroller"
	"github.com/autumnleaf-ra/golang/go-restapi-gin/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/product", productcontroller.Index)
	r.GET("/api/product/:id", productcontroller.Show)
	r.POST("/api/product", productcontroller.Create)
	r.PUT("/api/product/:id", productcontroller.Update)
	r.DELETE("/api/product/", productcontroller.Delete)

	r.Run()
}