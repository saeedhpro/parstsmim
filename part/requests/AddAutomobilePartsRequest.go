package requests

type AddAutomobilePartsRequest struct {
	Parts []AddAutomobilePart `json:"parts"`
}

type AddAutomobilePart struct {
	Name string `json:"name"`
}
