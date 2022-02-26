package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"part/logic"
)

type PartHandler interface {
	GetPart(c *gin.Context)
	GetAutomobileParts(c *gin.Context)
	AddFile(c *gin.Context)
	GetPartFiles(c *gin.Context)
	DownloadPartFiles(c *gin.Context)
	GetAutomobileFiles(c *gin.Context)
	AddAutomobileParts(c *gin.Context)
}

type partApi struct {
	getPartLogic            logic.GetPartLogic
	getAutomobilePartsLogic logic.GetAutomobilePartsLogic
	addFileLogic            logic.AddFileLogic
	getPartFilesLogic       logic.GetPartFilesLogic
	downloadPartFilesLogic  logic.DownloadPartFilesLogic
	getAutomobileFilesLogic logic.GetAutomobileFilesLogic
	addAutomobilePartsLogic logic.AddAutomobilePartsLogic
}

func NewPartApi(
	getPartLogic logic.GetPartLogic,
	getAutomobilePartsLogic logic.GetAutomobilePartsLogic,
	addFileLogic logic.AddFileLogic,
	getPartFilesLogic logic.GetPartFilesLogic,
	getAutomobileFilesLogic logic.GetAutomobileFilesLogic,
	addAutomobilePartsLogic logic.AddAutomobilePartsLogic,
	downloadPartFilesLogic logic.DownloadPartFilesLogic,
) *partApi {
	return &partApi{
		getPartLogic:            getPartLogic,
		getAutomobilePartsLogic: getAutomobilePartsLogic,
		addFileLogic:            addFileLogic,
		getPartFilesLogic:       getPartFilesLogic,
		downloadPartFilesLogic:  downloadPartFilesLogic,
		getAutomobileFilesLogic: getAutomobileFilesLogic,
		addAutomobilePartsLogic: addAutomobilePartsLogic,
	}
}

func (api *partApi) GetPart(c *gin.Context) {
	part, err := api.getPartLogic.GetPart(c)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(500, err.Error())
		return
	}
	if part == nil {
		fmt.Println("not found")
		c.JSON(404, "not found")
	}
	c.JSON(200, part)
	return
}

func (api *partApi) GetAutomobileParts(c *gin.Context) {
	parts, err := api.getAutomobilePartsLogic.GetAutomobileParts(c)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, parts)
	return
}

func (api *partApi) AddFile(c *gin.Context) {
	url, err := api.addFileLogic.AddFile(c)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(500, err.Error())
		return
	}
	if url == nil {
		fmt.Println("not found")
		c.JSON(404, "not found")
		return
	}
	c.JSON(200, url)
	return
}

func (api *partApi) GetPartFiles(c *gin.Context) {
	list := api.getPartFilesLogic.GetPartFiles(c)
	c.JSON(200, list)
	return
}

func (api *partApi) DownloadPartFiles(c *gin.Context) {
	path := api.downloadPartFilesLogic.DownloadPartFiles(c)
	if path == nil {
		fmt.Println("not found")
		c.JSON(404, "not found")
		return
	}
	c.File(*path)
	return
}

func (api *partApi) GetAutomobileFiles(c *gin.Context) {
	list, err := api.getAutomobileFilesLogic.GetAutomobileFiles(c)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, list)
	return
}

func (api *partApi) AddAutomobileParts(c *gin.Context) {
	api.addAutomobilePartsLogic.AddAutomobileParts(c)
}
