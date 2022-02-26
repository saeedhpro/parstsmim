package requests

type SendMailRequest struct {
	From    string   `json:"from"`
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Body    string   `json:"body"`
}

type EmailRequest struct {
	PartID string `json:"part_id"`
	File   string `json:"file"`
}
