package controllers

import (
	"net/http"
	"strconv"

	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/services"
	"github.com/gin-gonic/gin"
)

func ListEvents(c *gin.Context) {
	events, err := services.ListEvents()
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

	event, err := services.FindEvent(id)
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
