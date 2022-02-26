package logic

import (
	"automobile/helper"
	"automobile/model"
	"automobile/repository"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type GetAutomobileListLogic interface {
	GetAutomobileList(ctx *gin.Context) ([]model.Automobile, error)
}

type getAutomobileListLogic struct {
}

func NewGetAutomobileListLogic() GetAutomobileListLogic {
	return &getAutomobileListLogic{}
}

func (u *getAutomobileListLogic) GetAutomobileList(ctx *gin.Context) ([]model.Automobile, error) {
	list := []model.Automobile{}
	list, err := getAutomobileListByMySQL()
	if err != nil {
		return list, nil
	}
	return list, nil
}

func getAutomobileListByMySQL() ([]model.Automobile, error) {
	list := []model.Automobile{}
	query := "SELECT `id`, `model`, `manufacture`, `type` FROM `automobiles`"
	stmt, err := repository.DBS.MySQL.Prepare(query)
	if err != nil {
		log.Println(err.Error())
		return list, nil
	}
	rows, err := stmt.Query()
	if err != nil {
		log.Println(err.Error())
		return list, nil
	}
	var automobile model.Automobile
	for rows.Next() {
		err = rows.Scan(
			&automobile.ID,
			&automobile.Model,
			&automobile.Manufacture,
			&automobile.Type,
		)
		if err != nil {
			log.Println(err.Error())
			return list, nil
		}
		automobile.Parts = helper.GetParts(strconv.FormatInt(automobile.ID, 10))
		list = append(list, automobile)
	}
	return list, nil
}
