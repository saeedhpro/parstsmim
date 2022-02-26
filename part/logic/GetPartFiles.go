package logic

import (
	"github.com/gin-gonic/gin"
	"part/helper"
)

type GetPartFilesLogic interface {
	GetPartFiles(ctx *gin.Context) []string
}

type getPartFilesLogic struct {
}

func NewGetPartFilesLogic() GetPartFilesLogic {
	return &getPartFilesLogic{}
}

func (u *getPartFilesLogic) GetPartFiles(ctx *gin.Context) []string {
	files := []string{}
	id := ctx.Param("id")
	if id == "" {
		return files
	}
	return helper.GetPartFiles(id)
}
