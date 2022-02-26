package http

import (
	"file/logic"
	"fmt"

	"github.com/gin-gonic/gin"
)

type FileHandler interface {
	UploadFile(c *gin.Context)
	GetFile(c *gin.Context)
	SingleDownloadFile(c *gin.Context)
	GetFiles(c *gin.Context)
}

type fileApi struct {
	singleDownloadFile logic.DownloadSingleFileLogic
	uploadFileLogic    logic.UploadFileLogic
	getFileLogic       logic.GetFileLogic
	getFilesLogic      logic.GetFilesLogic
}

func NewFileApi(
	singleDownloadFile logic.DownloadSingleFileLogic,
	uploadFileLogic logic.UploadFileLogic,
	getFileLogic logic.GetFileLogic,
	getFilesLogic logic.GetFilesLogic,
) *fileApi {
	return &fileApi{
		singleDownloadFile: singleDownloadFile,
		uploadFileLogic:    uploadFileLogic,
		getFileLogic:       getFileLogic,
		getFilesLogic:      getFilesLogic,
	}
}

func (api *fileApi) UploadFile(c *gin.Context) {
	fileName, err := api.uploadFileLogic.UploadFile(c)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(200, fileName)
	return
}

func (api *fileApi) GetFile(c *gin.Context) {
	path, err := api.getFileLogic.GetFile(c)
	if err != nil {
		fmt.Println(err.Error())
		if path == nil {
			c.JSON(404, "not found")
			return
		}
		c.JSON(500, err.Error())
		return
	}
	c.File(*path)
	return
}

func (api *fileApi) SingleDownloadFile(c *gin.Context) {
	path, err := api.singleDownloadFile.DownloadSingleFile(c)
	if err != nil {
		fmt.Println(err.Error())
		if path == nil {
			c.JSON(404, "not found")
			return
		}
		c.JSON(500, err.Error())
		return
	}
	c.File(*path)
	return

}

func (api *fileApi) GetFiles(c *gin.Context) {
	list, err := api.getFilesLogic.GetFiles(c)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(200, list)
	return
}
