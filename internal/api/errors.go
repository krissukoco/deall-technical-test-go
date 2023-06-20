package api

import "errors"

var (
	ErrValidation = errors.New("validation error")
)

const (
	CodeInvalidJson        = 10000
	CodeInvalidCredentials = 10001
	CodeInvalidToken       = 10002
	CodeExpiredToken       = 10003
	CodeUnauthorized       = 10004
	CodeForbidden          = 10005
	CodeEmailAlreadyExists = 10010
	CodePasswordInvalid    = 10011
	CodeUserData           = 10012
	CodeUnknown            = 19999
)
