package model

type Automobile struct {
	ID          int64  `json:"id"`
	Model       int    `json:"model"`
	Manufacture string `json:"manufacture"`
	Type        string `json:"type"`
	Parts       []Part `json:"parts"`
}
