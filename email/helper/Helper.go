package helper

import (
	"email/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetAutomobile(id string) *model.Automobile {
	fmt.Println(id, "ema aut")
	response, err := http.Get(fmt.Sprintf("%s/automobiles/%s", "http://automobile:8001", id))
	if err != nil {
		fmt.Print(err.Error())
		return nil
	}
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	var automobile model.Automobile
	err = json.Unmarshal(responseData, &automobile)
	return &automobile
}

func GetPart(id string) *model.Part {
	fmt.Println(id, "pid")
	response, err := http.Get(fmt.Sprintf("%s/parts/%s", "http://part:8002", id))
	if err != nil {
		fmt.Print(err.Error())
		return nil
	}
	fmt.Println(response, "res")
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	var part model.Part
	err = json.Unmarshal(responseData, &part)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return &part
}
