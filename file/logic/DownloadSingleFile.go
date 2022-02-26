package logic

import (
	minio2 "file/minio"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"log"
)

type DownloadSingleFileLogic interface {
	DownloadSingleFile(ctx *gin.Context) (*string, error)
}

type downloadSingleFileLogic struct {
}

func NewDownloadSingleFileLogic() DownloadSingleFileLogic {
	return &downloadSingleFileLogic{}
}

func (u *downloadSingleFileLogic) DownloadSingleFile(ctx *gin.Context) (*string, error) {
	name := ctx.Param("name")
	if name == "" {
		return nil, nil
	}
	path := fmt.Sprintf("/tmp/%s", name)
	err := minio2.Client.FGetObject(ctx, minio2.BucketName, name, path, minio.GetObjectOptions{})
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return &path, nil
}
