package controller

import "part/logic"

var (
	GetPart            logic.GetPartLogic
	GetAutomobileParts logic.GetAutomobilePartsLogic
	AddFile            logic.AddFileLogic
	GetPartFiles       logic.GetPartFilesLogic
	DownloadPartFiles  logic.DownloadPartFilesLogic
	GetAutomobileFiles logic.GetAutomobileFilesLogic
	AddAutomobileParts logic.AddAutomobilePartsLogic
)

func init() {
	GetPart = logic.NewGetPartLogic()
	GetAutomobileParts = logic.NewGetAutomobilePartsLogic()
	AddFile = logic.NewAddFileLogic()
	GetPartFiles = logic.NewGetPartFilesLogic()
	DownloadPartFiles = logic.NewDownloadPartFilesLogic()
	GetAutomobileFiles = logic.NewGetAutomobileFilesLogic()
	AddAutomobileParts = logic.NewAddAutomobilePartsLogic()
}
