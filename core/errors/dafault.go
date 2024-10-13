package errors

import "net/http"

var (
	ErrUnAuthorized  = NewCustomHttpError(http.StatusUnauthorized, Unauthorized, "unauthorized")
	ErrTokenExpires  = NewCustomHttpError(http.StatusUnauthorized, Unauthorized, "token_expires")
	ErrPasswordWrong = NewCustomHttpError(http.StatusUnauthorized, PasswordWrong, "password_wrong")
)

var (
	ErrSystem           = NewCustomHttpError(http.StatusInternalServerError, SystemError, "system_error")
	ErrBadRequest       = NewCustomHttpError(http.StatusBadRequest, BadRequest, "invalid_request")
	ErrResourceNotFound = NewCustomHttpError(http.StatusNotFound, ResourceNotFound, "resource_not_found")
	ErrConflict         = NewCustomHttpError(http.StatusOK, Conflict, "username_has_exist ")
	ErrEmpty            = NewCustomHttpError(http.StatusOK, Success, "")
)
