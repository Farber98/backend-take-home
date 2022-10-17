package models

type Response struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func NewResponseWithoutData(msg string) Response {
	return Response{
		Message: msg,
	}
}

func NewResponseWithData(msg string, data interface{}) Response {
	return Response{
		Message: msg,
		Data:    data,
	}
}
