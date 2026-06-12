package loans

import (
	"time"

	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/entities"
)

type LoanResponse struct {
	ID         int64      `json:"id"`
	UserID     int64      `json:"user_id"`
	BookID     int64      `json:"book_id"`
	ReturnDate time.Time  `json:"return_date"`
	ReturnedAt *time.Time `json:"returned_at"`
	BookTitle  string     `json:"book_title"`
	BookAuthor string     `json:"book_author"`
}

func NewLoanResponse(l entities.Loan) LoanResponse {
	return LoanResponse{
		ID:         l.ID,
		UserID:     l.UserID,
		BookID:     l.BookID,
		ReturnDate: l.ReturnDate,
		ReturnedAt: l.ReturnedAt,
		BookTitle:  l.BookTitle,
		BookAuthor: l.BookAuthor,
	}
}

func NewLoanResponseList(list []entities.Loan) []LoanResponse {
	out := make([]LoanResponse, 0, len(list))
	for _, l := range list {
		out = append(out, NewLoanResponse(l))
	}
	return out
}
