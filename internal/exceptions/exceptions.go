package exceptions

import "errors"

var (
	ErrEmailNotFound        = errors.New("email not found")
	ErrEmailShouldBeUnique  = errors.New("email already registered")
	ErrInternalServerError  = errors.New("internal server error")
	ErrInvalidCredentials   = errors.New("invalid credentials")
	ErrNotAllowed           = errors.New("not allowed")
	ErrTimeoutExceeded      = errors.New("timeout exceeded")
	ErrUnauthorized         = errors.New("unauthorized")
	ErrUnproccessableEntity = errors.New("body could not be processed")
	ErrUserNotFound         = errors.New("user not found")
)
