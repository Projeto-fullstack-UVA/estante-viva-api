package services

import (
	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/models"
	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/repositories"
)

func ListBooks() ([]models.Book, error) {
	return repositories.GetBooks()
}

// FindBook returns the book with the given id, or ErrBookNotFound.
func FindBook(id int64) (*models.Book, error) {
	book, err := repositories.GetBookByID(id)
	if err != nil {
		return nil, err
	}
	if book == nil {
		return nil, ErrBookNotFound
	}
	return book, nil
}

// CreateBook inserts a new book, returning ErrBookCreateFailed when nothing was inserted.
func CreateBook(book models.Book) error {
	affected, err := repositories.CreateBook(book)
	if err != nil {
		return err
	}
	if affected == 0 {
		return ErrBookCreateFailed
	}
	return nil
}
