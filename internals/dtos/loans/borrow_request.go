package loans

import "time"

type BorrowRequest struct {
	BookID     int64     `json:"book_id" binding:"required"`
	ReturnDate time.Time `json:"return_date" binding:"required"`
}
