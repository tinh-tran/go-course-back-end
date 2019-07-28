package main

import (
	. "ImgService/common"
	"ImgService/services"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Static("/", "./public")
	v1 := router.Group("upload")
	{
		imgSer := new(services.ImageService)
		v1.POST("/category", imgSer.UploadCat)
		v1.POST("/course", imgSer.UploadCourse)
		videopdfSer := new(services.VideoPdfService)
		v1.POST("/video", videopdfSer.UploadVideo)
		v1.POST("/pdf", videopdfSer.UploadPdf)
	}

	port := GetServicePort()
	router.Run(":" + port)
}
