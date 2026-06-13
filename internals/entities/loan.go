package entities

import "time"

type Loan struct {
	ID         int64      `json:"id"`
	UserID     int64      `json:"user_id"`
	BookID     int64      `json:"book_id"`
	ReturnDate time.Time  `json:"return_date"`
	ReturnedAt *time.Time `json:"returned_at"`
	BookTitle  string     `json:"book_title"`
	BookAuthor string     `json:"book_author"`
	CreatedAt  time.Time  `json:"created_at"`
}
