package controllers

import (
	"errors"
	"net/http"
	"strconv"

	loandto "github.com/Projeto-fullstack-UVA/estante-viva-api/internals/dtos/loans"
	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/services"
	"github.com/gin-gonic/gin"
)

func ListLoans(c *gin.Context) {
	loans, err := services.ListLoans()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, loandto.NewLoanResponseList(loans))
}

func FindLoan(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid ID")
		return
	}

	loan, err := services.FindLoan(id)
	if err != nil {
		if errors.Is(err, services.ErrLoanNotFound) {
			c.String(http.StatusNotFound, "Loan not found")
			return
		}
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, loandto.NewLoanResponse(*loan))
}

func BorrowBook(c *gin.Context) {
	var req loandto.BorrowRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.String(http.StatusBadRequest, "Invalid loan payload — expected { user_id, book_id }")
		return
	}

	loan, err := services.BorrowBook(req.UserID, req.BookID)
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

	c.JSON(http.StatusCreated, loandto.NewLoanResponse(*loan))
}

func ReturnBook(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid ID")
		return
	}

	loan, err := services.ReturnBook(id)
	if err != nil {
		switch {
		case errors.Is(err, services.ErrLoanNotFound):
			c.String(http.StatusNotFound, "Loan not found")
		case errors.Is(err, services.ErrAlreadyReturned):
			c.String(http.StatusConflict, "Book already returned")
		default:
			c.String(http.StatusInternalServerError, "Error while returning book")
		}
		return
	}

	c.JSON(http.StatusOK, loandto.NewLoanResponse(*loan))
}
