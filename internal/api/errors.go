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

	// Subscriptions
	CodeNoSubscription          = 20000
	CodeSubscriptionRequired    = 20001
	CodeSubscriptionTypeInvalid = 20010
	CodeAlreadySubscribed       = 20011

	// Match
	CodeMatchNotFound     = 30001
	CodeMatchAlreadyLiked = 30002
	CodeNoMatchAvailable  = 30003

	// Internal
	CodeInternal = 90000
)
