package services

import (
	"context"
	"log"

	bookDto "github.com/Projeto-fullstack-UVA/estante-viva-api/internals/dtos/books"
	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/repositories"
)

func ListBooks(ctx context.Context) ([]bookDto.BookResponse, error) {
	books, err := repositories.GetBooks(ctx)
	if err != nil {
		log.Println("Error while fetching books from the database: ", err.Error())
		return nil, ErrListBooksFailed
	}

	log.Println("Success fetching all books from the database")

	return bookDto.NewBookResponseList(books), nil
}

func FindBook(ctx context.Context, id int64) (*bookDto.BookResponse, error) {
	book, err := repositories.GetBookByID(ctx, id)
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

func CreateBook(ctx context.Context, req bookDto.CreateBookRequest) error {
	affected, err := repositories.CreateBook(ctx, req.ToModel())
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

func DeleteBook(ctx context.Context, id int64) error {
	affected, err := repositories.DeleteBook(ctx, id)
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

func UpdateBook(ctx context.Context, id int64, req bookDto.UpdateBookRequest) error {
	book := req.ToModel()
	affected, err := repositories.UpdateBook(ctx, id, book)
	if err != nil {
		log.Println("Error while updating book in the database: ", err.Error())
		return ErrBookUpdateFailed
	}
	if affected == 0 {
		log.Println("No book with the id ", id, " found to update")
		return ErrBookNotFound
	}

	log.Println("Success updating book in the database")

	return nil
}
