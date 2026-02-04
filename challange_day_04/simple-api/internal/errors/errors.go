package errs

import "errors"

var (
	ErrInvalidInput      = errors.New("invalid input")
	ErrInvalidPagination = errors.New("invalid pagination parameters")
	ErrInvalidSort       = errors.New("invalid sort parameters")
	ErrInvalidSearch     = errors.New("search name cannot be empty")
	ErrInvalidEmail      = errors.New("email already exist")
	ErrInvalidPassword   = errors.New("invalid password")

	//when create
	ErrCreateUser = errors.New("error when create user")
	ErrAssignRole = errors.New("error when assigned user role")
	ErrTrxUser    = errors.New("transaction errors")
)
