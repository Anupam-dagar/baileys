package dto

type Status struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	Type       string `json:"type"`
	TotalCount *int   `json:"totalCount,omitempty"`
}

type BaseResponse struct {
	Status *Status     `json:"status"`
	Data   interface{} `json:"data"`
}
