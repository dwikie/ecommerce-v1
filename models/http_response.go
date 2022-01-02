package models

type HttpResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}

type HttpResponsePagination struct {
	Page       int64        `json:"page"`
	Limit      int64        `json:"limit"`
	TotalPages int64        `json:"total_page"`
	TotalData  int64        `json:"total_data"`
	Success    bool         `json:"success"`
	Data       *interface{} `json:"data"`
	Error      interface{}  `json:"error"`
}
