package entities

import (
	"time"
)

type Book struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Edition     string    `json:"edition"`
	ReleaseDate time.Time `json:"release_date"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}
