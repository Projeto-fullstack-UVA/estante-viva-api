package events

import (
	"time"

	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/entities"
)

type CreateEventRequest struct {
	Name          string    `json:"name" binding:"required"`
	Description   string    `json:"description" binding:"required"`
	Date          time.Time `json:"date" binding:"required"`
	Location      string    `json:"location" binding:"required"`
	InstitutionId int64     `json:"institution_id" binding:"required"`
}

func (r CreateEventRequest) ToModel() entities.Event {
	return entities.Event{
		Name:          r.Name,
		Description:   r.Description,
		Date:          r.Date,
		Location:      r.Location,
		InstitutionId: r.InstitutionId,
	}
}
