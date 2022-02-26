package http

import (
	"email/controller"
	"github.com/gin-gonic/gin"
)

func Run(port string) {
	engine := gin.Default()
	engine.Use(gin.Recovery())

	email := engine.Group("email")

	emailApi := NewMailApi(controller.SendMail)

	email.POST("/send", emailApi.SendMail)
	_ = engine.Run(port)
}
