package controllers

import (
	"net/http"

	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/services"
	"github.com/gin-gonic/gin"
)

func ListInstitutions(c *gin.Context) {
	institutions, err := services.ListInstitutions()
	if err != nil {
		c.String(http.StatusInternalServerError, "Error while fetching institutions: ")
		return
	}
	c.JSON(http.StatusOK, institutions)
}