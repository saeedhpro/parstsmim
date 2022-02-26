package logic

import (
	"github.com/gin-gonic/gin"
	"log"
	"part/repository"
)

type GetAutomobileFilesLogic interface {
	GetAutomobileFiles(ctx *gin.Context) ([]string, error)
}

type getAutomobileFilesLogic struct {
}

func NewGetAutomobileFilesLogic() GetAutomobileFilesLogic {
	return &getAutomobileFilesLogic{}
}

func (u *getAutomobileFilesLogic) GetAutomobileFiles(ctx *gin.Context) ([]string, error) {
	id := ctx.Param("id")
	if id == "" {
		return nil, nil
	}
	zipPath, err := getAutomobileFilesByMySQL(id)
	if err != nil {
		return nil, nil
	}
	return zipPath, nil
}

func getAutomobileFilesByMySQL(id string) ([]string, error) {
	files := []string{}
	query := "SELECT `part_files`.`name` FROM `part_files` LEFT JOIN `parts` on `part_files`.`part_id` = `parts`.`id` WHERE `parts`.`automobile_id` = ?"
	stmt, err := repository.DBS.MySQL.Prepare(query)
	if err != nil {
		log.Println(err.Error())
		return files, nil
	}
	rows, err := stmt.Query(id)
	if err != nil {
		log.Println(err.Error())
		return files, nil
	}
	var partFile string
	for rows.Next() {
		err = rows.Scan(
			&partFile,
		)
		if err != nil {
			log.Println(err.Error())
			return files, nil
		}
		files = append(files, partFile)
	}
	return files, nil
}
