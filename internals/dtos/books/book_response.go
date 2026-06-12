package books

import (
	"time"

	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/entities"
)

type BookResponse struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	ReleaseDate time.Time `json:"release_date"`
	Edition     string    `json:"edition"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewBookResponse(b entities.Book) BookResponse {
	return BookResponse{
		ID:          b.ID,
		Title:       b.Title,
		Author:      b.Author,
		ReleaseDate: b.ReleaseDate,
		Edition:     b.Edition,
		Status:      b.Status,
		CreatedAt:   b.CreatedAt,
	}
}

func NewBookResponseList(list []entities.Book) []BookResponse {
	out := make([]BookResponse, 0, len(list))
	for _, b := range list {
		out = append(out, NewBookResponse(b))
	}
	return out
}
