package logic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"part/helper"
	"part/repository"
	"part/requests"
)

type AddFileLogic interface {
	AddFile(ctx *gin.Context) (*string, error)
}

type addFileLogic struct {
}

func NewAddFileLogic() AddFileLogic {
	return &addFileLogic{}
}

func (u *addFileLogic) AddFile(ctx *gin.Context) (*string, error) {
	id := ctx.Param("id")
	part := helper.GetPart(id)
	if part == nil {
		fmt.Println("part nil")
		return nil, nil
	}
	var request requests.AddFileRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		fmt.Println("part nil")
		return nil, err
	}
	return addFileByMySQL(id, request.Name)
}

func addFileByMySQL(id string, name string) (*string, error) {
	query := "INSERT INTO `part_files`(`part_id`, `name`) VALUES (?,?)"
	stmt, err := repository.DBS.MySQL.Prepare(query)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	_, err = stmt.Exec(id, name)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	go sendEmail(id, name)
	return &name, nil
}

func sendEmail(id string, name string) {
	body := requests.SendMailRequestOnPartAdded{
		PartID: id,
		File:   name,
	}
	jsonData, err := json.Marshal(body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	http.Post(fmt.Sprintf("%s/email/send", "http://email:8004"), "application/json", bytes.NewBuffer(jsonData))
}
