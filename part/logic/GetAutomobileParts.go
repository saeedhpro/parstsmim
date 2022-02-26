package logic

import (
	"github.com/gin-gonic/gin"
	"log"
	"part/helper"
	"part/model"
	"part/repository"
	"strconv"
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
	parts, err := getAutomobilePartsByMySQL(id)
	if err != nil {
		return parts, nil
	}
	for i := 0; i < len(parts); i++ {
		files := helper.GetPartFiles(strconv.FormatInt(parts[i].ID, 10))
		parts[i].Files = files
	}
	return parts, nil
}

func getAutomobilePartsByMySQL(id string) ([]model.Part, error) {
	parts := []model.Part{}
	query := "SELECT id, name, automobile_id FROM `parts` where automobile_id = ?"
	stmt, err := repository.DBS.MySQL.Prepare(query)
	if err != nil {
		log.Println(err.Error())
		return parts, nil
	}
	rows, err := stmt.Query(id)
	if err != nil {
		log.Println(err.Error())
		return parts, nil
	}
	part := model.Part{}
	for rows.Next() {
		err = rows.Scan(
			&part.ID,
			&part.Name,
			&part.AutomobileID,
		)
		if err != nil {
			log.Println(err.Error())
			return parts, nil
		}
		parts = append(parts, part)
	}
	return parts, nil
}
