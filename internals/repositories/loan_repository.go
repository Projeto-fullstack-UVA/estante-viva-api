package repositories

import (
	"context"
	"errors"

	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/models"
	"github.com/jackc/pgx/v5"
)

const loanSelect = `SELECT l.id, l.user_id, l.book_id, l.return_date, l.returned_at,
	b.title AS book_title, b.author AS book_author
	FROM loans l
	JOIN books b ON l.book_id = b.id`

func scanLoan(row pgx.Row) (*models.Loan, error) {
	var l models.Loan
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

func GetLoans() ([]models.Loan, error) {
	rows, err := Pool.Query(context.Background(), loanSelect+" ORDER BY l.id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	loans := []models.Loan{}
	for rows.Next() {
		l, err := scanLoan(rows)
		if err != nil {
			return nil, err
		}
		loans = append(loans, *l)
	}
	return loans, rows.Err()
}

func GetLoanByID(id int64) (*models.Loan, error) {
	row := Pool.QueryRow(context.Background(), loanSelect+" WHERE l.id = $1", id)
	return scanLoan(row)
}

// CreateLoan inserts a loan and returns the id of the newly created row.
func CreateLoan(userID, bookID int64, returnDate string) (int64, error) {
	var id int64
	err := Pool.QueryRow(context.Background(),
		`INSERT INTO loans (user_id, book_id, return_date, returned_at)
		 VALUES ($1, $2, $3, NULL) RETURNING id`,
		userID, bookID, returnDate,
	).Scan(&id)
	return id, err
}

// ReturnLoan stamps returned_at and returns the number of rows affected.
func ReturnLoan(id int64, returnedAt string) (int64, error) {
	tag, err := Pool.Exec(context.Background(),
		`UPDATE loans SET returned_at = $1 WHERE id = $2`, returnedAt, id)
	if err != nil {
		return 0, err
	}
	return tag.RowsAffected(), nil
}
