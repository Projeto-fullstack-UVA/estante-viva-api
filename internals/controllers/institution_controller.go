package controllers

import (
	"errors"
	"net/http"
	"strconv"

	institutiondto "github.com/Projeto-fullstack-UVA/estante-viva-api/internals/dtos/institutions"
	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/services"
	"github.com/gin-gonic/gin"
)

func ListInstitutions(c *gin.Context) {
	institutions, err := services.ListInstitutions(c.Request.Context())
	if err != nil {
		c.String(http.StatusInternalServerError, "Error while fetching institutions")
		return
	}
	c.JSON(http.StatusOK, institutions)
}

func FindInstitution(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid id")
		return
	}

	institution, err := services.FindInstitution(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrInstitutionNotFound) {
			c.String(http.StatusNotFound, "Institution not found")
			return
		}
		c.String(http.StatusInternalServerError, "Error while fetching institution")
		return
	}
	c.JSON(http.StatusOK, institution)
}

func CreateInstitution(c *gin.Context) {
	var req institutiondto.CreateInstitutionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.String(http.StatusBadRequest, "Invalid institution format")
		return
	}

	if err := services.CreateInstitution(c.Request.Context(), req); err != nil {
		c.String(http.StatusInternalServerError, "Error while creating institution")
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Institution created successfully", "success": true})
}

func DeleteInstitution(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid id")
		return
	}

	if err := services.DeleteInstitution(c.Request.Context(), id); err != nil {
		if errors.Is(err, services.ErrInstitutionNotFound) {
			c.String(http.StatusNotFound, "Institution not found")
			return
		}
		c.String(http.StatusInternalServerError, "Error while deleting institution")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Institution deleted successfully", "success": true})
}
