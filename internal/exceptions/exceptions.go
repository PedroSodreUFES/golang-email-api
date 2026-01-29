package exceptions

import "errors"

var (
	ErrUnproccessableEntity = errors.New("body could not be processed")
	ErrInternalServerError  = errors.New("internal server error")
	ErrEmailShouldBeUnique  = errors.New("email already registered")
	ErrInvalidCredentials   = errors.New("invalid credentials")
	ErrUnauthorized         = errors.New("unauthorized")
	ErrUserNotFound         = errors.New("user not found")
	ErrTimeoutExceeded      = errors.New("timeout exceeded")
)
