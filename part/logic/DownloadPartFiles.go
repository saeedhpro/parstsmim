package logic

import (
	"github.com/gin-gonic/gin"
	"part/helper"
)

type DownloadPartFilesLogic interface {
	DownloadPartFiles(ctx *gin.Context) *string
}

type downloadPartFilesLogic struct {
}

func NewDownloadPartFilesLogic() DownloadPartFilesLogic {
	return &downloadPartFilesLogic{}
}

func (u *downloadPartFilesLogic) DownloadPartFiles(ctx *gin.Context) *string {
	id := ctx.Param("id")
	if id == "" {
		return nil
	}
	files := helper.GetPartFiles(id)
	zipPath := helper.GetZipFiles(files)
	return zipPath
}
