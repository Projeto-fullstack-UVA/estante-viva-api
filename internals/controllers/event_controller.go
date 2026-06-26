package controllers

import (
	"errors"
	"net/http"
	"strconv"

	eventdto "github.com/Projeto-fullstack-UVA/estante-viva-api/internals/dtos/events"
	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/services"
	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/utils"
	"github.com/gin-gonic/gin"
)

func ListEvents(c *gin.Context) {
	events, err := services.ListEvents(c.Request.Context())
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, "INTERNAL_ERROR", "Error while fetching events")
		return
	}
	utils.OK(c, http.StatusOK, events)
}

func FindEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.Fail(c, http.StatusBadRequest, "INVALID_ID", "Invalid ID")
		return
	}

	event, err := services.FindEvent(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrEventNotFound) {
			utils.Fail(c, http.StatusNotFound, "EVENT_NOT_FOUND", "Event not found")
			return
		}
		utils.Fail(c, http.StatusInternalServerError, "INTERNAL_ERROR", "Error while fetching event")
		return
	}
	utils.OK(c, http.StatusOK, event)
}

func CreateEvent(c *gin.Context) {
	var req eventdto.CreateEventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Fail(c, http.StatusBadRequest, "INVALID_PAYLOAD", "Invalid event format")
		return
	}

	if err := services.CreateEvent(c.Request.Context(), req); err != nil {
		utils.Fail(c, http.StatusInternalServerError, "INTERNAL_ERROR", "Failed to create event")
		return
	}

	utils.OK(c, http.StatusCreated, gin.H{"message": "Created event successfully"})
}

func DeleteEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.Fail(c, http.StatusBadRequest, "INVALID_ID", "Invalid ID")
		return
	}

	if err := services.DeleteEvent(c.Request.Context(), id); err != nil {
		utils.Fail(c, http.StatusInternalServerError, "INTERNAL_ERROR", "Failed to delete event")
		return
	}

	utils.OK(c, http.StatusOK, gin.H{"message": "Deleted event successfully"})
}
