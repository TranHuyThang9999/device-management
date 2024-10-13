package errors

import (
	"fmt"
	"net/http"
)

type Error interface {
	error
	GetCode() int
	GetMessage() string
	GetHttpCode() int
}

type CustomError struct {
	HttpCode int         `json:"http_code"`
	Code     int         `json:"code"`
	Message  string      `json:"message"`
	Body     interface{} `json:"body"`
}

func (c CustomError) GetCode() int {
	return c.Code
}

func (c CustomError) GetMessage() string {
	return c.Message
}

func (c CustomError) Error() string {
	return fmt.Sprintf("[%v] %s", c.Code, c.Message)
}

func (c CustomError) GetHttpCode() int {
	return c.HttpCode
}

func NewCustomHttpError(httCode int, code int, message string) *CustomError {
	return &CustomError{
		Code:     code,
		Message:  message,
		HttpCode: httCode,
	}
}

func NewCustomHttpErrorWithCode(code int, msg string, statusCode string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: fmt.Sprintf("%s with status code: [%v]", msg, statusCode),
	}
}

func NewSystemError(msg string) *CustomError {
	return &CustomError{
		HttpCode: http.StatusInternalServerError,
		Code:     SystemError,
		Message:  msg,
	}
}

func NewSystemErrorWithCode(statusCode string) *CustomError {
	return &CustomError{
		Code:    SystemError,
		Message: fmt.Sprintf("system error with status_code:[%v]", statusCode),
	}
}

func NewBadRequestWithCode(statusCode string) *CustomError {
	return &CustomError{
		HttpCode: http.StatusBadRequest,
		Code:     BadRequest,
		Message:  fmt.Sprintf("bad request with status_code:[%v]", statusCode),
	}
}

func NewUnknownError(statusCode string) *CustomError {
	msg := fmt.Sprintf("unknown error with code:[%v]", statusCode)
	return NewSystemError(msg)
}

func NewResourceNotFoundWithCode(statusCode string) *CustomError {
	return &CustomError{
		HttpCode: http.StatusNotFound,
		Code:     ResourceNotFound,
		Message:  fmt.Sprintf("resource not found with status_code:[%v]", statusCode),
	}
}

const (
	Success          = 0
	Unauthorized     = 1
	BadRequest       = 3
	SystemError      = 4
	Conflict         = 5
	ResourceNotFound = 6
	PasswordWrong    = 7
	UserStatusFalse  = 8
	Forbidden        = 10
)
