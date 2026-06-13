package services

import "errors"

var (
	ErrLoginWrongCredentials = errors.New("Wrong Credentials")

	ErrUserNotFound     = errors.New("User not found")
	ErrUserCreateFailed = errors.New("Error while creating user")

	ErrBookNotFound     = errors.New("Book not found")
	ErrBookCreateFailed = errors.New("Error while creating book")
	ErrBookNotAvailable = errors.New("Book is not available")

	ErrLoanNotFound     = errors.New("Loan not found")
	ErrLoanCreateFailed = errors.New("Error while creating loan")
	ErrAlreadyReturned  = errors.New("Book already returned")
	ErrReturnFailed     = errors.New("Error while returning book")
)
