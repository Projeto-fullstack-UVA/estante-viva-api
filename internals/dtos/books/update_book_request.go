package books

import (
	"time"

	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/entities"
)

type UpdateBookRequest struct {
	Title       *string    `json:"title"`
	Author      *string    `json:"author"`
	ReleaseDate *time.Time `json:"release_date"`
	Edition     *string    `json:"edition"`
	Status      *string    `json:"status" binding:"omitempty,oneof=available lent"`
}

func (r UpdateBookRequest) ToModel() entities.Book {
	book := entities.Book{}
	if r.Title != nil {
		book.Title = *r.Title
	}
	if r.Author != nil {
		book.Author = *r.Author
	}
	if r.ReleaseDate != nil {
		book.ReleaseDate = *r.ReleaseDate
	}
	if r.Edition != nil {
		book.Edition = *r.Edition
	}
	if r.Status != nil {
		book.Status = *r.Status
	}
	return book
}
