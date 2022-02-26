package http

import (
	"automobile/helper"
	"automobile/logic"
	"fmt"
	"github.com/gin-gonic/gin"
)

type AutomobileHandler interface {
	GetAutomobileParts(c *gin.Context)
	GetAutomobileList(c *gin.Context)
	GetAutomobile(c *gin.Context)
	GetAutomobileFiles(c *gin.Context)
	AddAutomobile(c *gin.Context)
}

type automobileApi struct {
	getAutomobilePartsLogic logic.GetAutomobilePartsLogic
	getAutomobileListLogic  logic.GetAutomobileListLogic
	getAutomobileLogic      logic.GetAutomobileLogic
	getAutomobileFilesLogic logic.GetAutomobileFilesLogic
	addAutomobile           logic.AddAutomobileLogic
}

func NewAutomobileApi(
	getAutomobilePartsLogic logic.GetAutomobilePartsLogic,
	getAutomobileListLogic logic.GetAutomobileListLogic,
	getAutomobileLogic logic.GetAutomobileLogic,
	getAutomobileFilesLogic logic.GetAutomobileFilesLogic,
	addAutomobile logic.AddAutomobileLogic,
) *automobileApi {
	return &automobileApi{
		getAutomobilePartsLogic: getAutomobilePartsLogic,
		getAutomobileListLogic:  getAutomobileListLogic,
		getAutomobileLogic:      getAutomobileLogic,
		getAutomobileFilesLogic: getAutomobileFilesLogic,
		addAutomobile:           addAutomobile,
	}
}

func (api *automobileApi) GetAutomobileList(c *gin.Context) {
	list, err := api.getAutomobileListLogic.GetAutomobileList(c)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(200, list)
	return
}

func (api *automobileApi) GetAutomobile(c *gin.Context) {
	automobile, err := api.getAutomobileLogic.GetAutomobile(c)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(500, automobile)
		return
	}
	if automobile == nil {
		fmt.Println("not found")
		c.JSON(404, "not found")
		return
	}
	c.JSON(200, automobile)
	return
}

func (api *automobileApi) GetAutomobileParts(c *gin.Context) {
	url, err := api.getAutomobilePartsLogic.GetAutomobileParts(c)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(200, url)
	return
}

func (api *automobileApi) GetAutomobileFiles(c *gin.Context) {
	list, err := api.getAutomobileFilesLogic.GetAutomobileFiles(c)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(500, err.Error())
		return
	}
	zipPath := helper.GetZipFiles(list)
	if zipPath != nil {
		c.File(*zipPath)
		return
	}
	c.JSON(400, "err")
	return
}

func (api *automobileApi) AddAutomobile(c *gin.Context) {
	list, err := api.addAutomobile.AddAutomobile(c)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, list)
	return
}
