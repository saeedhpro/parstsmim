package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"part/helper"
	"part/repository"
	"part/requests"
	"strings"
)

type AddAutomobilePartsLogic interface {
	AddAutomobileParts(ctx *gin.Context)
}

type addAutomobilePartsLogic struct {
}

func NewAddAutomobilePartsLogic() AddAutomobilePartsLogic {
	return &addAutomobilePartsLogic{}
}

func (u *addAutomobilePartsLogic) AddAutomobileParts(ctx *gin.Context) {
	id := ctx.Param("id")
	automobile := helper.GetAutomobile(id)
	if automobile == nil {
		ctx.JSON(404, "not found")
		return
	}
	var request requests.AddAutomobilePartsRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		fmt.Println(err.Error())
		ctx.JSON(500, err.Error())
		return
	}
	res, err := addAutomobilePartsByMySQL(id, request)
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(500, err.Error())
		return
	}
	if !res {
		fmt.Println("not found")
		ctx.JSON(404, "not found")
		return
	} else {
		ctx.JSON(200, "done")
		return
	}
}

func addAutomobilePartsByMySQL(id string, request requests.AddAutomobilePartsRequest) (bool, error) {
	automobile := helper.GetAutomobile(id)
	if automobile != nil {
		return addAutomobile(id, request.Parts)
	} else {
		return false, nil
	}
}

func addAutomobile(id string, parts []requests.AddAutomobilePart) (bool, error) {
	query := "INSERT INTO `parts`(`automobile_id`, `name`) VALUES "
	columns := []string{}
	var values []interface{}
	for _, r := range parts {
		columns = append(columns, " (?,?) ")
		values = append(values, id, r.Name)
	}
	columnsSrt := strings.Join(columns, ",")
	query += columnsSrt
	stmt, err := repository.DBS.MySQL.Prepare(query)
	if err != nil {
		log.Println(err.Error())
		return false, err
	}
	_, err = stmt.Exec(values...)
	if err != nil {
		log.Println(err.Error())
		return false, err
	}
	return true, nil
}
