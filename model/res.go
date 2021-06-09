package model

type Response struct {
	StatusCode int         `json:"status_code, omitempty"`
	Message    string      `json:"message, omitempty"`
	Data       interface{} `json:"data, omitempty"`
}
