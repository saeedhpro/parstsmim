package http

import (
	"github.com/gin-gonic/gin"
	"part/controller"
)

func Run(port string) {
	engine := gin.Default()
	engine.Use(gin.Recovery())

	parts := engine.Group("parts")

	partApi := NewPartApi(controller.GetPart, controller.GetAutomobileParts, controller.AddFile, controller.GetPartFiles, controller.GetAutomobileFiles, controller.AddAutomobileParts, controller.DownloadPartFiles)

	parts.GET("/:id", partApi.GetPart)
	parts.GET("/automobiles/:id/parts", partApi.GetAutomobileParts)
	parts.GET("/automobiles/:id/files", partApi.GetAutomobileFiles)
	parts.POST("/automobiles/:id/parts", partApi.AddAutomobileParts)
	parts.POST("/:id/file", partApi.AddFile)
	parts.GET("/:id/files", partApi.GetPartFiles)
	parts.GET("/:id/files/download", partApi.DownloadPartFiles)
	_ = engine.Run(port)
}
