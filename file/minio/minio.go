package minio

import (
	"fmt"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var Client *minio.Client

var BucketName = "pars"

func Run() {
	endPoint := "minio:9000"
	accessKeyID := os.Getenv("MINIO_ACCESS_KEY")
	secretAccessKey := os.Getenv("MINIO_SECRET_KEY")
	client, err := minio.New(endPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})
	Client = client
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%#v\n", Client)
}
