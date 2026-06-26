package controllers

import (
	"errors"
	"net/http"
	"strconv"

	institutiondto "github.com/Projeto-fullstack-UVA/estante-viva-api/internals/dtos/institutions"
	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/services"
	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/utils"
	"github.com/gin-gonic/gin"
)

func ListInstitutions(c *gin.Context) {
	institutions, err := services.ListInstitutions(c.Request.Context())
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, "INTERNAL_ERROR", "Error while fetching institutions")
		return
	}
	utils.OK(c, http.StatusOK, institutions)
}

func FindInstitution(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.Fail(c, http.StatusBadRequest, "INVALID_ID", "Invalid ID")
		return
	}

	institution, err := services.FindInstitution(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrInstitutionNotFound) {
			utils.Fail(c, http.StatusNotFound, "INSTITUTION_NOT_FOUND", "Institution not found")
			return
		}
		utils.Fail(c, http.StatusInternalServerError, "INTERNAL_ERROR", "Error while fetching institution")
		return
	}
	utils.OK(c, http.StatusOK, institution)
}

func CreateInstitution(c *gin.Context) {
	var req institutiondto.CreateInstitutionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Fail(c, http.StatusBadRequest, "INVALID_PAYLOAD", "Invalid institution format")
		return
	}

	if err := services.CreateInstitution(c.Request.Context(), req); err != nil {
		utils.Fail(c, http.StatusInternalServerError, "INTERNAL_ERROR", "Error while creating institution")
		return
	}

	utils.OK(c, http.StatusCreated, gin.H{"message": "Institution created successfully"})
}

func UpdateInstitution(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.Fail(c, http.StatusBadRequest, "INVALID_ID", "Invalid ID")
		return
	}

	var req institutiondto.UpdateInstitutionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Fail(c, http.StatusBadRequest, "INVALID_PAYLOAD", "Invalid institution format")
		return
	}

	if err := services.UpdateInstitution(c.Request.Context(), id, req); err != nil {
		if errors.Is(err, services.ErrInstitutionNotFound) {
			utils.Fail(c, http.StatusNotFound, "INSTITUTION_NOT_FOUND", "Institution not found")
			return
		}
		utils.Fail(c, http.StatusInternalServerError, "INTERNAL_ERROR", "Error while updating institution")
		return
	}

	utils.OK(c, http.StatusOK, gin.H{"message": "Institution updated successfully"})
}

func DeleteInstitution(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.Fail(c, http.StatusBadRequest, "INVALID_ID", "Invalid ID")
		return
	}

	if err := services.DeleteInstitution(c.Request.Context(), id); err != nil {
		if errors.Is(err, services.ErrInstitutionNotFound) {
			utils.Fail(c, http.StatusNotFound, "INSTITUTION_NOT_FOUND", "Institution not found")
			return
		}
		utils.Fail(c, http.StatusInternalServerError, "INTERNAL_ERROR", "Error while deleting institution")
		return
	}

	utils.OK(c, http.StatusOK, gin.H{"message": "Institution deleted successfully"})
}
