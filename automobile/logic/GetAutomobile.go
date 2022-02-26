package logic

import (
	"automobile/helper"
	"automobile/model"
	"github.com/gin-gonic/gin"
)

type GetAutomobileLogic interface {
	GetAutomobile(ctx *gin.Context) (*model.Automobile, error)
}

type getAutomobileLogic struct {
}

func NewGetAutomobileLogic() GetAutomobileLogic {
	return &getAutomobileLogic{}
}

func (u *getAutomobileLogic) GetAutomobile(ctx *gin.Context) (*model.Automobile, error) {
	id := ctx.Param("id")
	automobile, err := helper.GetAutomobile(id)
	if err != nil {
		return automobile, nil
	}
	if automobile == nil {
		return nil, nil
	}
	automobile.Parts = helper.GetParts(id)
	return automobile, nil
}
