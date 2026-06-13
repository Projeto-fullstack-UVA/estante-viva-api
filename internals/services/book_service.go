package services

import (
	"log"

	bookDto "github.com/Projeto-fullstack-UVA/estante-viva-api/internals/dtos/books"
	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/repositories"
)

func ListBooks() ([]bookDto.BookResponse, error) {
	books, err := repositories.GetBooks()
	if err != nil {
		log.Println("Error while fetching books from the database: ", err.Error())
		return nil, ErrListBooksFailed
	}

	log.Println("Success fetching all books from the database")

	return bookDto.NewBookResponseList(books), nil
}

func FindBook(id int64) (*bookDto.BookResponse, error) {
	book, err := repositories.GetBookByID(id)
	if err != nil {
		log.Println("Error while fetching book from the database: ", err.Error())
		return nil, ErrBookFetchFailed
	}
	if book == nil {
		log.Println("No book with the id ", id, " found in the database")
		return nil, ErrBookNotFound
	}

	resp := bookDto.NewBookResponse(*book)

	log.Println("Success fetching book from the database")

	return &resp, nil
}

func CreateBook(req bookDto.CreateBookRequest) error {
	affected, err := repositories.CreateBook(req.ToModel())
	if err != nil {
		log.Println("Error while creating book in the database: ", err.Error())
		return ErrBookCreateFailed
	}
	if affected == 0 {
		log.Println("Failed to register book")
		return ErrBookCreateFailed
	}

	log.Println("Success creating book in the database")

	return nil
}

func DeleteBook(id int64) error {
	affected, err := repositories.DeleteBook(id)
	if err != nil {
		log.Println("Error while deleting book from the database: ", err.Error())
		return ErrBookDeleteFailed
	}
	if affected == 0 {
		log.Println("No book with the id ", id, " found to delete")
		return ErrBookNotFound
	}

	log.Println("Success deleting book from the database")

	return nil
}
