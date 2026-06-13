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
		return nil, ErrListLoansFailed
	}

	log.Println("Book's list fetched from database with success")
	return loandto.NewLoanResponseList(loans), nil
}

func FindLoan(id int64) (*loandto.LoanResponse, error) {
	loan, err := repositories.GetLoanByID(id)
	if err != nil {
		log.Println("Error while fetching loan from the database: ", err.Error())
		return nil, ErrLoanFetchFailed
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
		log.Println("Error fetching book from the database: ", err)
		return nil, ErrBookFetchFailed
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
		return nil, ErrBookAlreadyReturned
	}

	returnedAt := time.Now()

	affected, err := repositories.ReturnLoan(id, returnedAt)
	if err != nil || affected == 0 {
		log.Println("Error while marking loan as returned in the database")
		return nil, ErrBookReturnFailed
	}

	if err := repositories.UpdateBookStatus(loan.BookID, "available"); err != nil {
		log.Println("Error while updating book status in the database: ", err.Error())
		return nil, ErrBookReturnFailed
	}

	return FindLoan(id)
}

func DeleteLoan(id int64) error {
	loan, err := repositories.GetLoanByID(id)
	if err != nil {
		log.Println("Error while fetching loan from the database: ", err.Error())
		return ErrLoanFetchFailed
	}
	if loan == nil {
		log.Println("Loan with the id ", id, " was not found in the database")
		return ErrLoanNotFound
	}

	affected, err := repositories.DeleteLoan(id)
	if err != nil {
		log.Println("Error while deleting loan from the database: ", err.Error())
		return ErrLoanDeleteFailed
	}
	if affected == 0 {
		log.Println("No loan with the id ", id, " found to delete")
		return ErrLoanNotFound
	}

	// Release the book if the loan was still active so it isn't stuck as "lent".
	if loan.ReturnedAt == nil {
		if err := repositories.UpdateBookStatus(loan.BookID, "available"); err != nil {
			log.Println("Error while releasing book after loan deletion: ", err.Error())
			return ErrLoanDeleteFailed
		}
	}

	log.Println("Success deleting loan from the database")

	return nil
}
