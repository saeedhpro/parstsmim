package http

import (
	"email/logic"
	"github.com/gin-gonic/gin"
)

type EmailHandler interface {
	SendMail(c *gin.Context)
}

type emailApi struct {
	sendMailLogic logic.SendMailLogic
}

func NewMailApi(
	sendMailLogic logic.SendMailLogic,
) *emailApi {
	return &emailApi{
		sendMailLogic: sendMailLogic,
	}
}

func (api *emailApi) SendMail(c *gin.Context) {
	api.sendMailLogic.SendMail(c)
}
