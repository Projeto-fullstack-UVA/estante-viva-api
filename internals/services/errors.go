package services

import "errors"

var (
	ErrUserNotFound     = errors.New("User not found")
	ErrUserCreateFailed = errors.New("Error while creating user")
	ErrUserFetchFailed  = errors.New("Error while fetching user")
	ErrUserDeleteFailed = errors.New("Error while deleting user")
	ErrUserUpdateFailed = errors.New("Error while updating user")
	ErrListUsersFailed  = errors.New("Error while listing users")

	ErrBookNotFound     = errors.New("Book not found")
	ErrBookCreateFailed = errors.New("Error while creating book")
	ErrBookFetchFailed  = errors.New("Error while fetching book")
	ErrBookDeleteFailed = errors.New("Error while deleting book")
	ErrBookUpdateFailed = errors.New("Error while updating book")
	ErrBookNotAvailable = errors.New("Book is not available")
	ErrListBooksFailed  = errors.New("Error while listing books")

	ErrLoanNotFound        = errors.New("Loan not found")
	ErrLoanCreateFailed    = errors.New("Error while creating loan")
	ErrLoanFetchFailed     = errors.New("Error while fetching loan")
	ErrLoanDeleteFailed    = errors.New("Error while deleting loan")
	ErrListLoansFailed     = errors.New("Error while listing loans")
	ErrBookAlreadyReturned = errors.New("Book already returned")
	ErrBookReturnFailed    = errors.New("Error while returning book")

	ErrInstitutionNotFound    = errors.New("Institution not found")
	ErrInstitutionFetchFailed = errors.New("Error while fetching institution")

	ErrEventNotFound     = errors.New("Event not found")
	ErrEventListFailed   = errors.New("Error while listing events")
	ErrCreateEventFailed = errors.New("Error while creating event")
	ErrDeleteEventFailed = errors.New("Error while deleting event")
)
