package logic

import (
	"automobile/helper"
	"automobile/model"
	"automobile/repository"
	"automobile/requests"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type AddAutomobileLogic interface {
	AddAutomobile(ctx *gin.Context) (*model.Automobile, error)
}

type addAutomobileLogic struct {
}

func NewAddAutomobileLogic() AddAutomobileLogic {
	return &addAutomobileLogic{}
}

func (u *addAutomobileLogic) AddAutomobile(ctx *gin.Context) (*model.Automobile, error) {
	var request requests.AddAutomobileRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return addAutomobileByMySQL(request)
}

func addAutomobileByMySQL(request requests.AddAutomobileRequest) (*model.Automobile, error) {
	query := "INSERT INTO `automobiles`(`type`, `manufacture`, `model`) VALUES (?,?,?)"
	stmt, err := repository.DBS.MySQL.Prepare(query)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	result, err := stmt.Exec(
		&request.Type,
		&request.Manufacture,
		&request.Model,
	)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	automobile, err := helper.GetAutomobile(strconv.FormatInt(id, 10))
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return automobile, nil
}
