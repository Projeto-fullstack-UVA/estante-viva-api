package controllers

import (
	"errors"
	"net/http"
	"strconv"

	loandto "github.com/Projeto-fullstack-UVA/estante-viva-api/internals/dtos/loans"
	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/middleware"
	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/services"
	"github.com/gin-gonic/gin"
)

func ListLoans(c *gin.Context) {
	loans, err := services.ListLoans(c.Request.Context())
	if err != nil {
		c.String(http.StatusInternalServerError, "Error while returning the loan's list")
		return
	}

	c.JSON(http.StatusOK, loans)
}

func FindLoan(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid ID")
		return
	}

	loan, err := services.FindLoan(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrLoanNotFound) {
			c.String(http.StatusNotFound, "Loan not found")
			return
		}
		c.String(http.StatusInternalServerError, "Error while finding loan")
		return
	}

	c.JSON(http.StatusOK, loan)
}

func BorrowBook(c *gin.Context) {
	var req loandto.BorrowRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.String(http.StatusBadRequest, "Invalid loan payload: expected { book_id }")
		return
	}

	userId, ok := middleware.GetCurrentUserID(c)
	if !ok {
		c.String(http.StatusUnauthorized, "Unauthorized")
		return
	}

	loan, err := services.BorrowBook(c.Request.Context(), userId, req.BookID, req.ReturnDate)
	if err != nil {
		switch {
		case errors.Is(err, services.ErrBookNotFound):
			c.String(http.StatusNotFound, "Book not found")
		case errors.Is(err, services.ErrBookNotAvailable):
			c.String(http.StatusConflict, "Book is not available")
		default:
			c.String(http.StatusInternalServerError, "Error while creating loan")
		}
		return
	}

	c.JSON(http.StatusCreated, loan)
}

func ReturnBook(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid ID")
		return
	}

	loan, err := services.ReturnBook(c.Request.Context(), id)
	if err != nil {
		switch {
		case errors.Is(err, services.ErrLoanNotFound):
			c.String(http.StatusNotFound, "Loan not found")
		case errors.Is(err, services.ErrBookAlreadyReturned):
			c.String(http.StatusConflict, "Book already returned")
		default:
			c.String(http.StatusInternalServerError, "Error while returning book")
		}
		return
	}

	c.JSON(http.StatusOK, loan)
}

func DeleteLoan(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := services.DeleteLoan(c.Request.Context(), id); err != nil {
		if errors.Is(err, services.ErrLoanNotFound) {
			c.String(http.StatusNotFound, "Loan not found")
			return
		}
		c.String(http.StatusInternalServerError, "Error while deleting loan")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Loan deleted successfully", "success": true})
}
