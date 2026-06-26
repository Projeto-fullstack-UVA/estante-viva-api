package controllers

import (
	"errors"
	"net/http"
	"strconv"

	loandto "github.com/Projeto-fullstack-UVA/estante-viva-api/internals/dtos/loans"
	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/middleware"
	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/services"
	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/utils"
	"github.com/gin-gonic/gin"
)

func ListLoans(c *gin.Context) {
	loans, err := services.ListLoans(c.Request.Context())
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, "INTERNAL_ERROR", "Error while returning the loan's list")
		return
	}

	utils.OK(c, http.StatusOK, loans)
}

func FindLoan(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.Fail(c, http.StatusBadRequest, "INVALID_ID", "Invalid ID")
		return
	}

	loan, err := services.FindLoan(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrLoanNotFound) {
			utils.Fail(c, http.StatusNotFound, "LOAN_NOT_FOUND", "Loan not found")
			return
		}
		utils.Fail(c, http.StatusInternalServerError, "INTERNAL_ERROR", "Error while finding loan")
		return
	}

	utils.OK(c, http.StatusOK, loan)
}

func BorrowBook(c *gin.Context) {
	var req loandto.BorrowRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Fail(c, http.StatusBadRequest, "INVALID_PAYLOAD", "Invalid loan payload: expected { book_id }")
		return
	}

	userId, ok := middleware.GetCurrentUserID(c)
	if !ok {
		utils.Fail(c, http.StatusUnauthorized, "UNAUTHORIZED", "Unauthorized")
		return
	}

	loan, err := services.BorrowBook(c.Request.Context(), userId, req.BookID, req.ReturnDate)
	if err != nil {
		switch {
		case errors.Is(err, services.ErrBookNotFound):
			utils.Fail(c, http.StatusNotFound, "BOOK_NOT_FOUND", "Book not found")
		case errors.Is(err, services.ErrBookNotAvailable):
			utils.Fail(c, http.StatusConflict, "BOOK_NOT_AVAILABLE", "Book is not available")
		default:
			utils.Fail(c, http.StatusInternalServerError, "INTERNAL_ERROR", "Error while creating loan")
		}
		return
	}

	utils.OK(c, http.StatusCreated, loan)
}

func ReturnBook(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.Fail(c, http.StatusBadRequest, "INVALID_ID", "Invalid ID")
		return
	}

	loan, err := services.ReturnBook(c.Request.Context(), id)
	if err != nil {
		switch {
		case errors.Is(err, services.ErrLoanNotFound):
			utils.Fail(c, http.StatusNotFound, "LOAN_NOT_FOUND", "Loan not found")
		case errors.Is(err, services.ErrBookAlreadyReturned):
			utils.Fail(c, http.StatusConflict, "BOOK_ALREADY_RETURNED", "Book already returned")
		default:
			utils.Fail(c, http.StatusInternalServerError, "INTERNAL_ERROR", "Error while returning book")
		}
		return
	}

	utils.OK(c, http.StatusOK, loan)
}

func DeleteLoan(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.Fail(c, http.StatusBadRequest, "INVALID_ID", "Invalid ID")
		return
	}

	if err := services.DeleteLoan(c.Request.Context(), id); err != nil {
		if errors.Is(err, services.ErrLoanNotFound) {
			utils.Fail(c, http.StatusNotFound, "LOAN_NOT_FOUND", "Loan not found")
			return
		}
		utils.Fail(c, http.StatusInternalServerError, "INTERNAL_ERROR", "Error while deleting loan")
		return
	}

	utils.OK(c, http.StatusOK, gin.H{"message": "Loan deleted successfully"})
}
