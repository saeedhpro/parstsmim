package helper

import (
	"archive/zip"
	"automobile/model"
	"automobile/repository"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func GetAutomobile(id string) (*model.Automobile, error) {
	automobile := model.Automobile{}
	query := "SELECT `id`, `model`, `manufacture`, `type` FROM `automobiles` WHERE `id` = ?"
	stmt, err := repository.DBS.MySQL.Prepare(query)
	if err != nil {
		log.Println(err.Error())
		return &automobile, nil
	}
	result := stmt.QueryRow(id)
	if err = result.Err(); err != nil {
		log.Println(err.Error())
		return nil, nil
	}
	err = result.Scan(
		&automobile.ID,
		&automobile.Model,
		&automobile.Manufacture,
		&automobile.Type,
	)
	if err != nil {
		log.Println(err.Error())
		return nil, nil
	}
	return &automobile, nil
}

func GetParts(id string) []model.Part {
	var parts []model.Part
	response, err := http.Get(fmt.Sprintf("%s/parts/automobiles/%s/parts", "http://part:8002", id))
	if err != nil {
		fmt.Print(err.Error())
		return parts
	}
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		return parts
	}
	err = json.Unmarshal(responseData, &parts)
	if err != nil {
		fmt.Println(err.Error())
		return parts
	}
	return parts
}

func GetAutomobileFiles(id string) []string {
	var files []string
	response, err := http.Get(fmt.Sprintf("%s/parts/automobiles/%s/files", "http://part:8002", id))
	if err != nil {
		fmt.Print(err.Error())
		return files
	}
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
		return files
	}
	err = json.Unmarshal(responseData, &files)
	if err != nil {
		fmt.Println(err.Error())
		return files
	}
	return files
}

func GetZipFiles(files []string) *string {
	var list []string
	var names []string
	var zipPath string
	for n, file := range files {
		if file == "" {
			continue
		}
		response, err := http.Get(fmt.Sprintf("%s/files/file/%s", "http://file:8003", file))
		if err != nil {
			fmt.Print(err.Error())
			return nil
		}
		name := fmt.Sprintf("%d-%s", n, file)
		path := fmt.Sprintf("/tmp/%s", name)
		out, err := os.Create(path)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		_, err = io.Copy(out, response.Body)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		defer out.Close()
		list = append(list, path)
		names = append(names, name)
	}
	zipPath = "/tmp/file.zip"
	zipFile, err := os.Create(zipPath)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	zipWriter := zip.NewWriter(zipFile)
	for i := 0; i < len(list); i++ {
		f, err := os.Open(list[i])
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		defer f.Close()
		w, err := zipWriter.Create(names[i])
		if err != nil {
			panic(err)
		}
		if _, err := io.Copy(w, f); err != nil {
			panic(err)
		}
	}
	zipWriter.Close()
	return &zipPath
}
