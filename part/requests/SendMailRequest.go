package requests

type SendMailRequestOnPartAdded struct {
	PartID string `json:"part_id"`
	File   string `json:"file"`
}
