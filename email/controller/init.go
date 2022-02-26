package controller

import "email/logic"

var SendMail logic.SendMailLogic

func init() {
	SendMail = logic.NewSendMailLogic()
}
