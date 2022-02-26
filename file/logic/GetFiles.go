package logic

import (
	minio2 "file/minio"
	"file/requests"
	"github.com/gin-gonic/gin"
	_ "github.com/minio/minio-go/v7"
	"log"
	"net/url"
	"time"
)

type GetFilesLogic interface {
	GetFiles(ctx *gin.Context) ([]string, error)
}

type getFilesLogic struct {
}

func NewGetFilesLogic() GetFilesLogic {
	return &getFilesLogic{}
}

func (u *getFilesLogic) GetFiles(ctx *gin.Context) ([]string, error) {
	var request requests.GetMultipleFileRequest
	if err := ctx.ShouldBindJSON(request); err != nil {
		log.Println(err.Error())
		return nil, err
	}
	var urlList []string
	reqParams := make(url.Values)
	for _, name := range request.NameList {
		preSignedURL, err := minio2.Client.PresignedGetObject(ctx, minio2.BucketName, name, time.Second*24*60*60, reqParams)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
		urlList = append(urlList, preSignedURL.RawPath)
	}
	return urlList, nil
}
