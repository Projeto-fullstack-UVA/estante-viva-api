package models

type Loan struct {
	ID         int64   `json:"id"`
	UserID     int64   `json:"user_id"`
	BookID     int64   `json:"book_id"`
	ReturnDate string  `json:"return_date"`
	ReturnedAt *string `json:"returned_at"`
	BookTitle  string  `json:"book_title"`
	BookAuthor string  `json:"book_author"`
}
