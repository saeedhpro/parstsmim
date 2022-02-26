package model

type Part struct {
	ID           int64    `json:"id" bson:"_id,omitempty"`
	AutomobileID int      `json:"automobile_id" bson:"automobile_id,omitempty"`
	Name         string   `json:"name" bson:"name,omitempty"`
	Files        []string `json:"files" bson:"files,omitempty"`
}
