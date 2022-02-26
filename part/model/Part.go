package model

type Part struct {
	ID           int64    `json:"id"`
	AutomobileID int64    `json:"automobile_id"`
	Name         string   `json:"name"`
	Files        []string `json:"files"`
}
