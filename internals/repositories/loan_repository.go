package repositories

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/entities"
	"github.com/jackc/pgx/v5"
)

var (
	ErrLoanNotFound        = errors.New("loan not found")
	ErrLoanAlreadyReturned = errors.New("loan already returned")
)

const loanSelect = `SELECT l.id, l.user_id, l.book_id, l.return_date, l.returned_at,
	b.title AS book_title, b.author AS book_author
	FROM loans l
	JOIN books b ON l.book_id = b.id`

func scanLoan(row pgx.Row) (*entities.Loan, error) {
	var l entities.Loan
	err := row.Scan(
		&l.ID, &l.UserID, &l.BookID, &l.ReturnDate,
		&l.ReturnedAt, &l.BookTitle, &l.BookAuthor,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &l, nil
}

func GetLoans(ctx context.Context) ([]entities.Loan, error) {
	rows, err := Pool.Query(ctx, loanSelect+" ORDER BY l.id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	loans := []entities.Loan{}
	for rows.Next() {
		l, err := scanLoan(rows)
		if err != nil {
			return nil, err
		}
		loans = append(loans, *l)
	}
	return loans, rows.Err()
}

func GetLoanByID(ctx context.Context, id int64) (*entities.Loan, error) {
	row := Pool.QueryRow(ctx, loanSelect+" WHERE l.id = $1", id)
	return scanLoan(row)
}

func CreateLoan(ctx context.Context, userID, bookID int64, returnDate time.Time) (*int64, error) {
	var id int64

	log.Println("Beginning transaction to borrow book...")

	tx, err := Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		log.Println("Failed to begin transaction", err)
		return nil, err
	}

	log.Println("Transaction started")

	defer tx.Rollback(ctx)

	log.Println("Creating loan's record...")
	err = tx.QueryRow(ctx,
		`INSERT INTO loans (user_id, book_id, return_date, returned_at)
		 VALUES ($1, $2, $3, NULL) RETURNING id`,
		userID, bookID, returnDate,
	).Scan(&id)
	if err != nil {
		log.Println("Failed to create loan's record:", err)
		return nil, err
	}
	if id == 0 {
		log.Println("The book is not available for borrowing")
		return nil, errors.New("No rows were affected")
	}
	log.Println("Created loan's record successfully")

	log.Println("Updating book's status...")
	result, err := tx.Exec(ctx,
		`UPDATE books SET status = $1 WHERE id = $2 AND status = 'available'`,
		"lent", bookID)
	if err != nil {
		log.Println("Failed to update book's status:", err)
		return nil, err
	}
	if result.RowsAffected() == 0 {
		log.Println("Book not available for borrowing")
		return nil, errors.New("Book not available for borrowing")
	}
	log.Println("Updated book's status successfully")

	log.Println("Committing transaction...")
	if err := tx.Commit(ctx); err != nil {
		log.Println("Failed to commit transaction:", err)
		return nil, err
	}
	log.Println("Committed transaction with success")

	log.Println("New loan's id: ", id)
	return &id, nil
}

func ReturnLoan(ctx context.Context, loanID int64) (*int64, error) {
	var l entities.Loan

	log.Printf("Beginning transaction to return book...")
	tx, err := Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		log.Printf("Failed to begin transaction: %v", err)
		return nil, err
	}
	log.Printf("Transaction started")

	defer tx.Rollback(ctx)

	log.Printf("Getting loan by id %v...", loanID)
	err = tx.QueryRow(ctx,
		`SELECT id, book_id, returned_at FROM loans WHERE id = $1`, loanID).Scan(&l.ID, &l.BookID, &l.ReturnedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.Printf("Loan with the id %v was not found. Aborting transaction.", loanID)
			return nil, ErrLoanNotFound
		}
		log.Printf("Failed to get loan: %v. Aborting transaction.", err)
		return nil, err
	}
	if l.ReturnedAt != nil {
		log.Printf("This loan was already finished, the book was returned at %v. Aborting transaction", l.ReturnedAt)
		return nil, ErrLoanAlreadyReturned
	}

	log.Printf("Updating loan's returned date...")
	_, err = tx.Exec(ctx,
		`UPDATE loans SET returned_at = $1 WHERE id = $2`, time.Now(), l.ID)
	if err != nil {
		log.Printf("Error while updating loan's returned date: %v", err)
		return nil, err
	}
	log.Printf("Updated loan's returned date successfully")

	log.Printf("Updating book's status'...")
	_, err = tx.Exec(ctx, `UPDATE books SET status = $1 WHERE id = $2`, "available", l.BookID)
	if err != nil {
		log.Printf("Failed to update book's status: %v \nAborting transaction.", err)
		return nil, err
	}
	log.Printf("Updated book's status' successfully.")

	log.Printf("Committing transaction...")
	if err := tx.Commit(ctx); err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		return nil, err
	}
	log.Println("Committed transaction with success")

	log.Printf("The book was returned successfully")
	return &l.ID, nil
}

func DeleteLoan(ctx context.Context, id int64) (int64, error) {
	var l entities.Loan

	log.Printf("Beginning transaction to delete loan...")
	tx, err := Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		log.Printf("Failed to begin transaction: %v", err)
		return 0, err
	}
	log.Printf("Transaction started")

	defer tx.Rollback(ctx)

	log.Printf("Getting loan by id %v...", id)
	err = tx.QueryRow(ctx,
		`SELECT id, book_id, returned_at FROM loans WHERE id = $1`, id).Scan(&l.ID, &l.BookID, &l.ReturnedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.Printf("Loan with the id %v was not found. Aborting transaction.", id)
			return 0, ErrLoanNotFound
		}
		log.Printf("Failed to get loan: %v. Aborting transaction.", err)
		return 0, err
	}

	log.Printf("Deleting loan's record...")
	tag, err := tx.Exec(ctx, `DELETE FROM loans WHERE id = $1`, id)
	if err != nil {
		log.Printf("Failed to delete loan: %v", err)
		return 0, err
	}

	// Release the book if the loan was still active so it isn't stuck as "lent".
	if l.ReturnedAt == nil {
		log.Printf("Loan was active, releasing book %v...", l.BookID)
		if _, err := tx.Exec(ctx,
			`UPDATE books SET status = $1 WHERE id = $2`, "available", l.BookID); err != nil {
			log.Printf("Failed to release book: %v \nAborting transaction.", err)
			return 0, err
		}
	}

	log.Printf("Committing transaction...")
	if err := tx.Commit(ctx); err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		return 0, err
	}
	log.Printf("Committed transaction with success")

	return tag.RowsAffected(), nil
}
