package http

import (
	"file/controller"

	"github.com/gin-gonic/gin"
)

func Run(port string) {
	engine := gin.Default()
	engine.Use(gin.Recovery())

	files := engine.Group("files")

	fileApi := NewFileApi(controller.SingleDownloadFile, controller.UploadFile, controller.GetFile, controller.GetFiles)

	files.POST("/", fileApi.UploadFile)
	files.GET("/file/:name", fileApi.GetFile)
	files.GET("/file/:name/download", fileApi.SingleDownloadFile)
	files.POST("/files", fileApi.GetFiles)
	_ = engine.Run(port)
}
