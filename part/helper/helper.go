package helper

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"part/model"
	"part/repository"
)

func GetAutomobile(id string) *model.Automobile {
	response, err := http.Get(fmt.Sprintf("%s/automobiles/%s", "http://automobile:8001", id))
	if err != nil {
		fmt.Print(err.Error())
		return nil
	}
	defer response.Body.Close()
	if response.StatusCode == 404 {
		fmt.Println(response.StatusCode)
		return nil
	}
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
	var part model.Part
	query := "SELECT `id`, `name`, `automobile_id` FROM `parts` WHERE `id` = ?"
	stmt, err := repository.DBS.MySQL.Prepare(query)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	res := stmt.QueryRow(id)
	if res.Err() != nil {
		fmt.Println(err.Error())
		return nil
	}
	err = res.Scan(
		&part.ID,
		&part.Name,
		&part.AutomobileID,
	)
	if err != nil {
		fmt.Println(err.Error())
		return &part
	}
	files := GetPartFiles(id)
	part.Files = files
	return &part
}

func GetPartFiles(id string) []string {
	files := []string{}
	query := "SELECT `part_files`.`name` FROM `part_files` WHERE `part_files`.`part_id` = ?"
	stmt, err := repository.DBS.MySQL.Prepare(query)
	if err != nil {
		log.Println(err.Error())
		return files
	}
	rows, err := stmt.Query(id)
	if err != nil {
		log.Println(err.Error())
		return files
	}
	var name string
	for rows.Next() {
		err = rows.Scan(
			&name,
		)
		if err != nil {
			log.Println(err.Error())
			return files
		}
		files = append(files, name)
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
