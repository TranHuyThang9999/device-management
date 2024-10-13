package resource

import "device_management/core/errors"

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Body    interface{} `json:"body,omitempty"`
}

func NewResponseResource(code int, message string, body interface{}) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Body:    body,
	}
}

func NewResponseErr(err errors.Error) *Response {
	return &Response{
		Code:    err.GetCode(),
		Message: err.GetMessage(),
		Body:    nil,
	}
}
