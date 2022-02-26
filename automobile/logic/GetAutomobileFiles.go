package logic

import (
	"automobile/helper"
	"fmt"
	"github.com/gin-gonic/gin"
)

type GetAutomobileFilesLogic interface {
	GetAutomobileFiles(ctx *gin.Context) ([]string, error)
}

type getAutomobileFilesLogic struct {
}

func NewGetAutomobileFilesLogic() GetAutomobileFilesLogic {
	return &getAutomobileFilesLogic{}
}

func (u *getAutomobileFilesLogic) GetAutomobileFiles(ctx *gin.Context) ([]string, error) {
	id := ctx.Param("id")
	files := []string{}
	if id == "" {
		return files, nil
	}
	automobile, err := helper.GetAutomobile(id)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	if automobile == nil {
		fmt.Println("automobile not found")
		return nil, nil
	}
	files = helper.GetAutomobileFiles(id)
	if err != nil {
		return files, nil
	}
	return files, nil
}
