package services

import (
	"context"
	"errors"
	"log"
	"time"

	loandto "github.com/Projeto-fullstack-UVA/estante-viva-api/internals/dtos/loans"
	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/repositories"
)

func ListLoans(ctx context.Context) ([]loandto.LoanResponse, error) {
	loans, err := repositories.GetLoans(ctx)
	if err != nil {
		log.Println("Error while fetching loans from database", err.Error())
		return nil, ErrListLoansFailed
	}

	log.Println("Book's list fetched from database with success")
	return loandto.NewLoanResponseList(loans), nil
}

func FindLoan(ctx context.Context, id int64) (*loandto.LoanResponse, error) {
	loan, err := repositories.GetLoanByID(ctx, id)
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

func BorrowBook(ctx context.Context, userID, bookID int64, returnDate time.Time) (*loandto.LoanResponse, error) {
	id, err := repositories.CreateLoan(ctx, userID, bookID, returnDate)
	if err != nil {
		log.Println("Error creating loan register in the database: ", err.Error())
		return nil, ErrLoanCreateFailed
	}
	log.Println("Book borrowed with success")

	return FindLoan(ctx, *id)
}

func ReturnBook(ctx context.Context, id int64) (*loandto.LoanResponse, error) {
	loanID, err := repositories.ReturnLoan(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, repositories.ErrLoanNotFound):
			return nil, ErrLoanNotFound
		case errors.Is(err, repositories.ErrLoanAlreadyReturned):
			return nil, ErrBookAlreadyReturned
		default:
			log.Printf("Error while marking loan as returned in the database: %v", err)
			return nil, ErrBookReturnFailed
		}
	}

	return FindLoan(ctx, *loanID)
}

func DeleteLoan(ctx context.Context, id int64) error {
	affected, err := repositories.DeleteLoan(ctx, id)
	if err != nil {
		if errors.Is(err, repositories.ErrLoanNotFound) {
			log.Println("Loan with the id ", id, " was not found in the database")
			return ErrLoanNotFound
		}
		log.Println("Error while deleting loan from the database: ", err.Error())
		return ErrLoanDeleteFailed
	}
	if affected == 0 {
		log.Println("No loan with the id ", id, " found to delete")
		return ErrLoanNotFound
	}

	log.Println("Success deleting loan from the database")

	return nil
}
