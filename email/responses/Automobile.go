package responses

type AutomobileList struct {
	List []Automobile `json:"list"`
}

type Automobile struct {
	ID          int64  `json:"id"`
	Model       int64  `json:"model"`
	Manufacture string `json:"manufacture"`
	Type        string `json:"type"`
}
