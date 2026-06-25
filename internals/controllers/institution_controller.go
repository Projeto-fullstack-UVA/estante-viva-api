package controllers

import (
	"errors"
	"net/http"
	"strconv"

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
		}
		c.String(http.StatusInternalServerError, "Error while fetching institution")
	}
	c.JSON(http.StatusOK, institution)
}
