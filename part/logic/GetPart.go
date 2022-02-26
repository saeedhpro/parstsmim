package logic

import (
	"github.com/gin-gonic/gin"
	"log"
	"part/helper"
	"part/model"
	"part/repository"
	"strconv"
)

type GetPartLogic interface {
	GetPart(ctx *gin.Context) (*model.Part, error)
}

type getPartLogic struct {
}

func NewGetPartLogic() GetPartLogic {
	return &getPartLogic{}
}

func (u *getPartLogic) GetPart(ctx *gin.Context) (*model.Part, error) {
	id := ctx.Param("id")
	if id == "" {
		return nil, nil
	}
	part, err := getPartByMySQL(id)
	if err != nil {
		return nil, nil
	}
	files := helper.GetPartFiles(strconv.FormatInt(part.ID, 10))
	part.Files = files
	return part, nil
}

func getPartByMySQL(id string) (*model.Part, error) {
	var part model.Part
	query := "SELECT id, name, automobile_id FROM `parts` where id = ?"
	stmt, err := repository.DBS.MySQL.Prepare(query)
	if err != nil {
		log.Println(err.Error())
		return nil, nil
	}
	result := stmt.QueryRow(id)
	err = result.Scan(
		&part.ID,
		&part.Name,
		&part.AutomobileID,
	)
	if err != nil {
		log.Println(err.Error())
		return nil, nil
	}
	return &part, nil
}
