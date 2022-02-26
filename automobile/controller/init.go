package controller

import "automobile/logic"

var (
	GetAutomobileParts logic.GetAutomobilePartsLogic
	GetAutomobileList  logic.GetAutomobileListLogic
	GetAutomobile      logic.GetAutomobileLogic
	GetAutomobileFiles logic.GetAutomobileFilesLogic
	AddAutomobile      logic.AddAutomobileLogic
)

func init() {
	GetAutomobileParts = logic.NewGetAutomobilePartsLogic()
	GetAutomobileList = logic.NewGetAutomobileListLogic()
	GetAutomobile = logic.NewGetAutomobileLogic()
	GetAutomobileFiles = logic.NewGetAutomobileFilesLogic()
	AddAutomobile = logic.NewAddAutomobileLogic()
}
