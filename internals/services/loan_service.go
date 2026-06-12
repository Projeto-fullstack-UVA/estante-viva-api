package services

import (
	"time"

	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/models"
	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/repositories"
)

// isoTimestamp mirrors JavaScript's Date.toISOString() (UTC, millisecond precision).
const isoTimestamp = "2006-01-02T15:04:05.000Z"

func ListLoans() ([]models.Loan, error) {
	return repositories.GetLoans()
}

// FindLoan returns the loan with the given id, or ErrLoanNotFound.
func FindLoan(id int64) (*models.Loan, error) {
	loan, err := repositories.GetLoanByID(id)
	if err != nil {
		return nil, err
	}
	if loan == nil {
		return nil, ErrLoanNotFound
	}
	return loan, nil
}

// BorrowBook lends an available book to a user and marks it as lent. The
// return date is set to 14 days from now.
func BorrowBook(userID, bookID int64) (*models.Loan, error) {
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

	return repositories.GetLoanByID(id)
}

// ReturnBook marks an outstanding loan as returned and frees up the book.
func ReturnBook(id int64) (*models.Loan, error) {
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

	return repositories.GetLoanByID(id)
}
