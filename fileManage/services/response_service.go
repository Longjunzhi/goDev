package services

type Response struct {
	Code    int64       `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func NewResponse() *Response {
	return &Response{}
}

func NewSuccessResponse() *Response {
	return &Response{Code: 0, Message: "success"}
}

func NewErrorResponse() *Response {
	return &Response{Code: 1, Message: "系统异常"}
}
