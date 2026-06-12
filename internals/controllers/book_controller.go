package controllers

import (
	"errors"
	"net/http"
	"strconv"

	bookdto "github.com/Projeto-fullstack-UVA/estante-viva-api/internals/dtos/books"
	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/services"
	"github.com/gin-gonic/gin"
)

func ListBooks(c *gin.Context) {
	books, err := services.ListBooks()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, bookdto.NewBookResponseList(books))
}

func FindBook(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid ID")
		return
	}

	book, err := services.FindBook(id)
	if err != nil {
		if errors.Is(err, services.ErrBookNotFound) {
			c.String(http.StatusNotFound, "Book not found")
			return
		}
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, bookdto.NewBookResponse(*book))
}

func CreateBook(c *gin.Context) {
	var req bookdto.CreateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.String(http.StatusBadRequest, "Invalid book format")
		return
	}

	if err := services.CreateBook(req.ToModel()); err != nil {
		c.String(http.StatusInternalServerError, "Error while creating book")
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Book created successfully", "success": true})
}
