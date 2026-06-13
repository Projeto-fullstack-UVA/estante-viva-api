package repositories

import (
	"context"
	"errors"
	"strconv"

	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/entities"
	"github.com/jackc/pgx/v5"
)

func scanBook(row pgx.Row) (*entities.Book, error) {
	var (
		id      int64
		edition *string
		b       entities.Book
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

func GetBooks() ([]entities.Book, error) {
	rows, err := Pool.Query(context.Background(),
		`SELECT id, title, author, release_date, edition, status, created_at FROM books ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := []entities.Book{}
	for rows.Next() {
		b, err := scanBook(rows)
		if err != nil {
			return nil, err
		}
		books = append(books, *b)
	}
	return books, rows.Err()
}

func GetBookByID(id int64) (*entities.Book, error) {
	row := Pool.QueryRow(context.Background(),
		`SELECT id, title, author, release_date, edition, status, created_at FROM books WHERE id = $1`, id)
	return scanBook(row)
}

func CreateBook(book entities.Book) (int64, error) {
	var edition *string
	if book.Edition != "" {
		edition = &book.Edition
	}
	tag, err := Pool.Exec(context.Background(),
		`INSERT INTO books (title, author, release_date, edition, status, created_at)
		 VALUES ($1, $2, $3, $4, $5, $6)`,
		book.Title, book.Author, book.ReleaseDate, edition, book.Status, book.CreatedAt,
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

func DeleteBook(id int64) (int64, error) {
	tag, err := Pool.Exec(context.Background(), `DELETE FROM books WHERE id = $1`, id)
	if err != nil {
		return 0, err
	}
	return tag.RowsAffected(), nil
}
