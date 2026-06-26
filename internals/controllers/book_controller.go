package controllers

import (
	"errors"
	"net/http"
	"strconv"

	bookdto "github.com/Projeto-fullstack-UVA/estante-viva-api/internals/dtos/books"
	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/services"
	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/utils"
	"github.com/gin-gonic/gin"
)

func ListBooks(c *gin.Context) {
	books, err := services.ListBooks(c.Request.Context())
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, "INTERNAL_ERROR", "Error while returning the book's list")
		return
	}

	utils.OK(c, http.StatusOK, books)
}

func FindBook(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.Fail(c, http.StatusBadRequest, "INVALID_ID", "Invalid ID")
		return
	}

	book, err := services.FindBook(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrBookNotFound) {
			utils.Fail(c, http.StatusNotFound, "BOOK_NOT_FOUND", "Book not found")
			return
		}
		utils.Fail(c, http.StatusInternalServerError, "INTERNAL_ERROR", "Error while finding book")
		return
	}

	utils.OK(c, http.StatusOK, book)
}

func CreateBook(c *gin.Context) {
	var req bookdto.CreateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Fail(c, http.StatusBadRequest, "INVALID_PAYLOAD", "Invalid book format")
		return
	}

	if err := services.CreateBook(c.Request.Context(), req); err != nil {
		utils.Fail(c, http.StatusInternalServerError, "INTERNAL_ERROR", "Error while creating book")
		return
	}

	utils.OK(c, http.StatusCreated, gin.H{"message": "Book created successfully"})
}

func DeleteBook(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.Fail(c, http.StatusBadRequest, "INVALID_ID", "Invalid ID")
		return
	}

	if err := services.DeleteBook(c.Request.Context(), id); err != nil {
		if errors.Is(err, services.ErrBookNotFound) {
			utils.Fail(c, http.StatusNotFound, "BOOK_NOT_FOUND", "Book not found")
			return
		}
		utils.Fail(c, http.StatusInternalServerError, "INTERNAL_ERROR", "Error while deleting book")
		return
	}

	utils.OK(c, http.StatusOK, gin.H{"message": "Book deleted successfully"})
}

func UpdateBook(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.Fail(c, http.StatusBadRequest, "INVALID_ID", "Invalid ID")
		return
	}

	var req bookdto.UpdateBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Fail(c, http.StatusBadRequest, "INVALID_PAYLOAD", "Invalid book format")
		return
	}

	if err := services.UpdateBook(c.Request.Context(), id, req); err != nil {
		if errors.Is(err, services.ErrBookNotFound) {
			utils.Fail(c, http.StatusNotFound, "BOOK_NOT_FOUND", "Book not found")
			return
		}
		utils.Fail(c, http.StatusInternalServerError, "INTERNAL_ERROR", "Error while updating book")
		return
	}

	utils.OK(c, http.StatusOK, gin.H{"message": "Book updated successfully"})
}
