package books

import (
	"time"

	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/models"
)

// BookResponse is the book representation returned to clients.
type BookResponse struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	ReleaseDate string    `json:"release_date"`
	Edition     string    `json:"edition"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewBookResponse(b models.Book) BookResponse {
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

func NewBookResponseList(list []models.Book) []BookResponse {
	out := make([]BookResponse, 0, len(list))
	for _, b := range list {
		out = append(out, NewBookResponse(b))
	}
	return out
}
