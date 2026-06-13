package services

import (
	"log"
	"time"

	loandto "github.com/Projeto-fullstack-UVA/estante-viva-api/internals/dtos/loans"
	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/repositories"
)

func ListLoans() ([]loandto.LoanResponse, error) {
	loans, err := repositories.GetLoans()
	if err != nil {
		log.Println("Error while fetching loans from database", err.Error())
		return nil, err
	}

	log.Println("Book's list fetched from database with success")
	return loandto.NewLoanResponseList(loans), nil
}

func FindLoan(id int64) (*loandto.LoanResponse, error) {
	loan, err := repositories.GetLoanByID(id)
	if err != nil {
		log.Println("Error while fetching loan from the database: ", err.Error())
		return nil, err
	}
	if loan == nil {
		log.Println("Loan with the id ", id, " was not found in the database")
		return nil, ErrLoanNotFound
	}

	resp := loandto.NewLoanResponse(*loan)

	log.Println("Loan found with success")

	return &resp, nil
}

func BorrowBook(userID, bookID int64) (*loandto.LoanResponse, error) {
	book, err := repositories.GetBookByID(bookID)
	if err != nil {
		log.Println("Error updating book status in the database: ", err)
		return nil, ErrReturnFailed
	}
	if book == nil {
		log.Println("No book with the id ", bookID, " was found in the database")
		return nil, ErrBookNotFound
	}
	if book.Status != "available" {
		log.Println("This book is not available for borrowing")
		return nil, ErrBookNotAvailable
	}

	returnDate := time.Now().AddDate(0, 0, 14)

	id, err := repositories.CreateLoan(userID, bookID, returnDate)
	if err != nil {
		log.Println("Error creating loan register in the database: ", err.Error())
		return nil, ErrLoanCreateFailed
	}

	if err := repositories.UpdateBookStatus(bookID, "lent"); err != nil {
		log.Println("Error updating book's status in the database: ", err)
		return nil, ErrLoanCreateFailed
	}

	log.Println("Book borrowed with success")

	return FindLoan(id)
}

func ReturnBook(id int64) (*loandto.LoanResponse, error) {
	loan, err := repositories.GetLoanByID(id)
	if err != nil {
		log.Println("Failed to fetch book from the database")
		return nil, ErrLoanNotFound
	}
	if loan == nil {
		log.Println("Loan with the id ", id, " was not found")
		return nil, ErrLoanNotFound
	}
	if loan.ReturnedAt != nil {
		log.Println("This loan was already finished, the book was returned at ", loan.ReturnedAt)
		return nil, ErrAlreadyReturned
	}

	returnedAt := time.Now()

	affected, err := repositories.ReturnLoan(id, returnedAt)
	if err != nil || affected == 0 {
		return nil, ErrReturnFailed
	}

	if err := repositories.UpdateBookStatus(loan.BookID, "available"); err != nil {
		return nil, err
	}

	return FindLoan(id)
}
