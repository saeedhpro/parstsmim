package requests

type AddAutomobileRequest struct {
	Type        string `json:"type"`
	Manufacture string `json:"manufacture"`
	Model       int    `json:"model"`
}
