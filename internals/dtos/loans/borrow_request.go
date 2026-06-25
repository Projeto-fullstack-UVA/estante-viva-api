package loans

type BorrowRequest struct {
	BookID int64 `json:"book_id" binding:"required"`
}
