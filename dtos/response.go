package dtos

type Response struct {
	Code  int         `json:"code,omitempty"`
	Data  interface{} `json:"data,omitempty"`
	Error error       `json:"error,omitempty"`
}
