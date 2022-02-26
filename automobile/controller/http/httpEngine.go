package http

import (
	"automobile/controller"
	"github.com/gin-gonic/gin"
)

func Run(port string) {
	engine := gin.Default()
	engine.Use(gin.Recovery())

	automobiles := engine.Group("automobiles")

	automobileApi := NewAutomobileApi(controller.GetAutomobileParts, controller.GetAutomobileList, controller.GetAutomobile, controller.GetAutomobileFiles, controller.AddAutomobile)

	automobiles.POST("/", automobileApi.AddAutomobile)
	automobiles.GET("/", automobileApi.GetAutomobileList)
	automobiles.GET("/:id", automobileApi.GetAutomobile)
	automobiles.GET("/:id/parts", automobileApi.GetAutomobileParts)
	automobiles.GET("/:id/files", automobileApi.GetAutomobileFiles)
	_ = engine.Run(port)
}
