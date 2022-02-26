package controller

import "file/logic"

var (
	UploadFile         logic.UploadFileLogic
	GetFile            logic.GetFileLogic
	SingleDownloadFile logic.DownloadSingleFileLogic
	GetFiles           logic.GetFilesLogic
)

func init() {
	UploadFile = logic.NewUploadFileLogic()
	GetFile = logic.NewGetFileLogic()
	SingleDownloadFile = logic.NewDownloadSingleFileLogic()
	GetFiles = logic.NewGetFilesLogic()
}
