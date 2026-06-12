package repositories

import (
	"context"
	"errors"
	"strconv"

	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/models"
	"github.com/jackc/pgx/v5"
)

func scanBook(row pgx.Row) (*models.Book, error) {
	var (
		id      int64
		edition *string
		b       models.Book
	)
	err := row.Scan(&id, &b.Title, &b.Author, &b.ReleaseDate, &edition, &b.Status, &b.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	b.ID = strconv.FormatInt(id, 10)
	if edition != nil {
		b.Edition = *edition
	}
	return &b, nil
}

func GetBooks() ([]models.Book, error) {
	rows, err := Pool.Query(context.Background(),
		`SELECT id, title, author, release_date, edition, status, created_at FROM books ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := []models.Book{}
	for rows.Next() {
		b, err := scanBook(rows)
		if err != nil {
			return nil, err
		}
		books = append(books, *b)
	}
	return books, rows.Err()
}

func GetBookByID(id int64) (*models.Book, error) {
	row := Pool.QueryRow(context.Background(),
		`SELECT id, title, author, release_date, edition, status, created_at FROM books WHERE id = $1`, id)
	return scanBook(row)
}

// CreateBook inserts a book and returns the number of rows affected.
func CreateBook(b models.Book) (int64, error) {
	var edition *string
	if b.Edition != "" {
		edition = &b.Edition
	}
	tag, err := Pool.Exec(context.Background(),
		`INSERT INTO books (title, author, release_date, edition, status, created_at)
		 VALUES ($1, $2, $3, $4, $5, $6)`,
		b.Title, b.Author, b.ReleaseDate, edition, b.Status, b.CreatedAt,
	)
	if err != nil {
		return 0, err
	}
	return tag.RowsAffected(), nil
}

func UpdateBookStatus(id int64, status string) error {
	_, err := Pool.Exec(context.Background(), `UPDATE books SET status = $1 WHERE id = $2`, status, id)
	return err
}
