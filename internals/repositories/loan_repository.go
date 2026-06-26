package repositories

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/entities"
	"github.com/jackc/pgx/v5"
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

func ReturnLoan(ctx context.Context, id int64) (int64, error) {
	tag, err := Pool.Exec(ctx,
		`UPDATE loans SET returned_at = $1 WHERE id = $2`, time.Now(), id)
	if err != nil {
		return 0, err
	}
	return tag.RowsAffected(), nil
}

func DeleteLoan(ctx context.Context, id int64) (int64, error) {
	tag, err := Pool.Exec(ctx, `DELETE FROM loans WHERE id = $1`, id)
	if err != nil {
		return 0, err
	}
	return tag.RowsAffected(), nil
}
