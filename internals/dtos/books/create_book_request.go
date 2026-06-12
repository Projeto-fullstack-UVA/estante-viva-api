package books

import (
	"time"

	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/entities"
)

// CreateBookRequest is the expected body for POST /books.
type CreateBookRequest struct {
	Title       string    `json:"title" binding:"required"`
	Author      string    `json:"author" binding:"required"`
	ReleaseDate time.Time `json:"release_date" binding:"required"`
	Edition     string    `json:"edition"`
	Status      string    `json:"status" binding:"required,oneof=available lent"`
	CreatedAt   time.Time `json:"created_at" binding:"required"`
}

func (r CreateBookRequest) ToModel() entities.Book {
	return entities.Book{
		Title:       r.Title,
		Author:      r.Author,
		ReleaseDate: r.ReleaseDate,
		Edition:     r.Edition,
		Status:      r.Status,
		CreatedAt:   r.CreatedAt,
	}
}
