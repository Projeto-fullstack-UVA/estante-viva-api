package services

import (
	bookDto "github.com/Projeto-fullstack-UVA/estante-viva-api/internals/dtos/books"
	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/repositories"
)

func ListBooks() ([]bookDto.BookResponse, error) {
	books, err := repositories.GetBooks()
	if err != nil {
		return nil, err
	}
	return bookDto.NewBookResponseList(books), nil
}

// FindBook returns the book with the given id, or ErrBookNotFound.
func FindBook(id int64) (*bookDto.BookResponse, error) {
	book, err := repositories.GetBookByID(id)
	if err != nil {
		return nil, err
	}
	if book == nil {
		return nil, ErrBookNotFound
	}
	resp := bookDto.NewBookResponse(*book)
	return &resp, nil
}

func CreateBook(req bookDto.CreateBookRequest) error {
	affected, err := repositories.CreateBook(req.ToModel())
	if err != nil {
		return err
	}
	if affected == 0 {
		return ErrBookCreateFailed
	}
	return nil
}
