package loans

// BorrowRequest is the expected body for POST /loans.
type BorrowRequest struct {
	UserID int64 `json:"user_id" binding:"required"`
	BookID int64 `json:"book_id" binding:"required"`
}
