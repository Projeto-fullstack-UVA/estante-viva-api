package services

import (
	"time"

	loandto "github.com/Projeto-fullstack-UVA/estante-viva-api/internals/dtos/loans"
	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/repositories"
)

// isoTimestamp mirrors JavaScript's Date.toISOString() (UTC, millisecond precision).
const isoTimestamp = "2006-01-02T15:04:05.000Z"

func ListLoans() ([]loandto.LoanResponse, error) {
	loans, err := repositories.GetLoans()
	if err != nil {
		return nil, err
	}
	return loandto.NewLoanResponseList(loans), nil
}

func FindLoan(id int64) (*loandto.LoanResponse, error) {
	loan, err := repositories.GetLoanByID(id)
	if err != nil {
		return nil, err
	}
	if loan == nil {
		return nil, ErrLoanNotFound
	}
	resp := loandto.NewLoanResponse(*loan)
	return &resp, nil
}

func BorrowBook(userID, bookID int64) (*loandto.LoanResponse, error) {
	book, err := repositories.GetBookByID(bookID)
	if err != nil {
		return nil, err
	}
	if book == nil {
		return nil, ErrBookNotFound
	}
	if book.Status != "available" {
		return nil, ErrBookNotAvailable
	}

	returnDate := time.Now().UTC().AddDate(0, 0, 14).Format(isoTimestamp)

	id, err := repositories.CreateLoan(userID, bookID, returnDate)
	if err != nil {
		return nil, ErrLoanCreateFailed
	}

	if err := repositories.UpdateBookStatus(bookID, "lent"); err != nil {
		return nil, err
	}

	return FindLoan(id)
}

func ReturnBook(id int64) (*loandto.LoanResponse, error) {
	loan, err := repositories.GetLoanByID(id)
	if err != nil {
		return nil, err
	}
	if loan == nil {
		return nil, ErrLoanNotFound
	}
	if loan.ReturnedAt != nil {
		return nil, ErrAlreadyReturned
	}

	returnedAt := time.Now().UTC().Format(isoTimestamp)

	affected, err := repositories.ReturnLoan(id, returnedAt)
	if err != nil || affected == 0 {
		return nil, ErrReturnFailed
	}

	if err := repositories.UpdateBookStatus(loan.BookID, "available"); err != nil {
		return nil, err
	}

	return FindLoan(id)
}
