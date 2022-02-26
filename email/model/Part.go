package model

type Part struct {
	ID           int    `json:"id,omitempty"`
	AutomobileID int    `json:"automobile_id,omitempty"`
	Name         string `json:"name,omitempty"`
}
