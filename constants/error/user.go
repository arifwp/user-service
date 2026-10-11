package error

import "errors"

var (
	ErrUserNotFound          = errors.New("user not found")
	ErrPasswordInCorrect     = errors.New("password incorrect")
	ErrUsernameAlreadyExists = errors.New("username already exists")
	ErrEmailAlreadyExists    = errors.New("email already exists")
	ErrPasswordDoesNotMatch  = errors.New("password does not match")
)

var UserErrors = []error{
	ErrUserNotFound,
	ErrPasswordInCorrect,
	ErrUsernameAlreadyExists,
	ErrEmailAlreadyExists,
	ErrPasswordDoesNotMatch,
}
