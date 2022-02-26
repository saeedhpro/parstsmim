package helper

import (
	"archive/zip"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"
)

const charset = "0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func RandomString(length int) string {
	return StringWithCharset(length, charset)
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
