package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/entities"
	"github.com/jackc/pgx/v5"
)

func scanBook(row pgx.Row) (*entities.Book, error) {
	var (
		edition *string
		b       entities.Book
	)
	err := row.Scan(&b.ID, &b.Title, &b.Author, &b.ReleaseDate, &edition, &b.Status, &b.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	if edition != nil {
		b.Edition = *edition
	}
	return &b, nil
}

func GetBooks(ctx context.Context) ([]entities.Book, error) {
	rows, err := Pool.Query(ctx,
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

func GetBookByID(ctx context.Context, id int64) (*entities.Book, error) {
	row := Pool.QueryRow(ctx,
		`SELECT id, title, author, release_date, edition, status, created_at FROM books WHERE id = $1`, id)
	return scanBook(row)
}

func CreateBook(ctx context.Context, book entities.Book) (int64, error) {
	var edition *string
	if book.Edition != "" {
		edition = &book.Edition
	}
	tag, err := Pool.Exec(ctx,
		`INSERT INTO books (title, author, release_date, edition, status, created_at)
		 VALUES ($1, $2, $3, $4, $5, $6)`,
		book.Title, book.Author, book.ReleaseDate, edition, book.Status, time.Now(),
	)
	if err != nil {
		return 0, err
	}
	return tag.RowsAffected(), nil
}

func UpdateBookStatus(ctx context.Context, id int64, status string) error {
	_, err := Pool.Exec(ctx, `UPDATE books SET status = $1 WHERE id = $2`, status, id)
	return err
}

func UpdateBook(ctx context.Context, id int64, book entities.Book) (int64, error) {
	var edition *string
	if book.Edition != "" {
		edition = &book.Edition
	}

	tag, err := Pool.Exec(ctx,
		`UPDATE books SET title = COALESCE(NULLIF($1, ''), title),
		 author = COALESCE(NULLIF($2, ''), author),
		 release_date = COALESCE(NULLIF($3, '0001-01-01'::timestamp), release_date),
		 edition = COALESCE(NULLIF($4::text, ''), edition),
		 status = COALESCE(NULLIF($5, ''), status)
		 WHERE id = $6`,
		book.Title, book.Author, book.ReleaseDate, edition, book.Status, id,
	)
	if err != nil {
		return 0, err
	}
	return tag.RowsAffected(), nil
}

func DeleteBook(ctx context.Context, id int64) (int64, error) {
	tag, err := Pool.Exec(ctx, `DELETE FROM books WHERE id = $1`, id)
	if err != nil {
		return 0, err
	}
	return tag.RowsAffected(), nil
}
