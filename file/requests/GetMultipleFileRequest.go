package requests

type GetMultipleFileRequest struct {
	NameList []string `json:"name_list"`
}
