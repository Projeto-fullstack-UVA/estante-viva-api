package books

import (
	"time"

	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/entities"
)

type CreateBookRequest struct {
	Title       string    `json:"title" binding:"required"`
	Author      string    `json:"author" binding:"required"`
	ReleaseDate time.Time `json:"release_date" binding:"required"`
	Edition     string    `json:"edition"`
	Status      string    `json:"status" binding:"required,oneof=available lent"`
}

func (r CreateBookRequest) ToModel() entities.Book {
	return entities.Book{
		Title:       r.Title,
		Author:      r.Author,
		ReleaseDate: r.ReleaseDate,
		Edition:     r.Edition,
		Status:      r.Status,
	}
}
