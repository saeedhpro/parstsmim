package model

type Automobile struct {
	ID          int    `json:"id,omitempty"`
	Model       int    `json:"model,omitempty"`
	Type        string `json:"type,omitempty"`
	Manufacture string `json:"manufacture,omitempty"`
	Parts       []Part `json:"parts"`
}
