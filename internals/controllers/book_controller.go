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
		c.String(http.StatusInternalServerError, "Error while returning the book's list")
		return
	}

	c.JSON(http.StatusOK, books)
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
		c.String(http.StatusInternalServerError, "Error while finding book")
		return
	}

	c.JSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context) {
	var req bookdto.CreateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.String(http.StatusBadRequest, "Invalid book format")
		return
	}

	if err := services.CreateBook(req); err != nil {
		c.String(http.StatusInternalServerError, "Error while creating book")
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Book created successfully", "success": true})
}

func DeleteBook(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := services.DeleteBook(id); err != nil {
		if errors.Is(err, services.ErrBookNotFound) {
			c.String(http.StatusNotFound, "Book not found")
			return
		}
		c.String(http.StatusInternalServerError, "Error while deleting book")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully", "success": true})
}

func UpdateBook(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid ID")
		return
	}

	var req bookdto.UpdateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.String(http.StatusBadRequest, "Invalid book format")
		return
	}

	if err := services.UpdateBook(id, req); err != nil {
		if errors.Is(err, services.ErrBookNotFound) {
			c.String(http.StatusNotFound, "Book not found")
			return
		}
		c.String(http.StatusInternalServerError, "Error while updating book")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully", "success": true})
}
