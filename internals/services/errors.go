package services

import "errors"

var (
	ErrUserNotFound     = errors.New("user not found")
	ErrUserCreateFailed = errors.New("error while creating user")

	ErrBookNotFound     = errors.New("book not found")
	ErrBookCreateFailed = errors.New("error while creating book")
	ErrBookNotAvailable = errors.New("book is not available")

	ErrLoanNotFound     = errors.New("loan not found")
	ErrLoanCreateFailed = errors.New("error while creating loan")
	ErrAlreadyReturned  = errors.New("book already returned")
	ErrReturnFailed     = errors.New("error while returning book")
)
