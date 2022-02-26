package logic

import (
	"file/helper"
	minio2 "file/minio"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

type UploadFileLogic interface {
	UploadFile(ctx *gin.Context) (*string, error)
}

type uploadFileLogic struct {
}

func NewUploadFileLogic() UploadFileLogic {
	return &uploadFileLogic{}
}

func (u *uploadFileLogic) UploadFile(ctx *gin.Context) (*string, error) {
	file, err := ctx.FormFile("file")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	buffer, err := file.Open()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	objectName := fmt.Sprintf("%s-%s", helper.RandomString(32), file.Filename)
	contentType := file.Header.Get("Content-Type")
	fileSize := file.Size
	_, err = minio2.Client.PutObject(ctx, minio2.BucketName, objectName, buffer, fileSize, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return &objectName, nil
}

func uploadFile(ctx *gin.Context, file *multipart.FileHeader) (*string, error) {
	t := time.Now().UnixNano()
	fileName := fmt.Sprintf("%d%s", t, filepath.Ext(file.Filename))
	path := "./files"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, 0755)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
	}
	err := ctx.SaveUploadedFile(file, path+"/"+fileName)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	//path = fmt.Sprintf("/files/%s", fileName)
	path = fmt.Sprintf("http://%s/files/%s", ctx.Request.Host, fileName)
	return &path, nil
}
