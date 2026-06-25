package controllers

import (
	"net/http"
	"strconv"

	eventdto "github.com/Projeto-fullstack-UVA/estante-viva-api/internals/dtos/events"
	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/services"
	"github.com/gin-gonic/gin"
)

func ListEvents(c *gin.Context) {
	events, err := services.ListEvents(c.Request.Context())
	if err != nil {
		c.String(http.StatusInternalServerError, "Error while fetching events")
		return
	}
	c.JSON(http.StatusOK, events)
}

func FindEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid id")
		return
	}

	event, err := services.FindEvent(c.Request.Context(), id)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error while fetching event")
		return
	}
	if event == nil {
		c.String(http.StatusNotFound, "Event not found")
		return
	}
	c.JSON(http.StatusOK, event)
}

func CreateEvent(c *gin.Context) {
	var req eventdto.CreateEventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.String(http.StatusBadRequest, "Invalid event format")
		return
	}

	if err := services.CreateEvent(c.Request.Context(), req); err != nil {
		c.String(http.StatusInternalServerError, "Failed to create event")
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Created event successfully", "success": true})
}

func DeleteEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid id")
		return
	}
	
	if err := services.DeleteEvent(c.Request.Context(), id); err != nil {
		c.String(http.StatusInternalServerError, "Failed to delete event")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted event successfully", "success": false})
}
