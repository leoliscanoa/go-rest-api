package dto

type ResponseDto[T any] struct {
	Time          string `json:"time"`
	Message       string `json:"message"`
	Data          T      `json:"data,omitempty"`
	Error         string `json:"error,omitempty"`
	TotalElements int    `json:"totalElements,omitempty"`
	Page          int    `json:"page,omitempty"`
	Size          int    `json:"size,omitempty"`
}
