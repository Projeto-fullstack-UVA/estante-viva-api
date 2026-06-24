package services

import "errors"

var (
	ErrUserNotFound     = errors.New("user not found")
	ErrUserCreateFailed = errors.New("error while creating user")
	ErrUserFetchFailed  = errors.New("error while fetching user")
	ErrUserDeleteFailed = errors.New("error while deleting user")
	ErrUserUpdateFailed = errors.New("error while updating user")
	ErrListUsersFailed  = errors.New("error while listing users")

	ErrBookNotFound     = errors.New("book not found")
	ErrBookCreateFailed = errors.New("error while creating book")
	ErrBookFetchFailed  = errors.New("error while fetching book")
	ErrBookDeleteFailed = errors.New("error while deleting book")
	ErrBookUpdateFailed = errors.New("error while updating book")
	ErrBookNotAvailable = errors.New("book is not available")
	ErrListBooksFailed  = errors.New("error while listing books")

	ErrLoanNotFound        = errors.New("loan not found")
	ErrLoanCreateFailed    = errors.New("error while creating loan")
	ErrLoanFetchFailed     = errors.New("error while fetching loan")
	ErrLoanDeleteFailed    = errors.New("error while deleting loan")
	ErrListLoansFailed     = errors.New("error while listing loans")
	ErrBookAlreadyReturned = errors.New("book already returned")
	ErrBookReturnFailed    = errors.New("error while returning book")
)
