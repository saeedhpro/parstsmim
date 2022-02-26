package logic

import (
	"automobile/helper"
	"automobile/model"
	"github.com/gin-gonic/gin"
)

type GetAutomobilePartsLogic interface {
	GetAutomobileParts(ctx *gin.Context) ([]model.Part, error)
}

type getAutomobilePartsLogic struct {
}

func NewGetAutomobilePartsLogic() GetAutomobilePartsLogic {
	return &getAutomobilePartsLogic{}
}

func (u *getAutomobilePartsLogic) GetAutomobileParts(ctx *gin.Context) ([]model.Part, error) {
	id := ctx.Param("id")
	parts := []model.Part{}
	if id == "" {
		return parts, nil
	}
	parts = helper.GetParts(id)
	return parts, nil
}
