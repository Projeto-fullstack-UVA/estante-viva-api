package events

import (
	"time"

	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/entities"
)

type UpdateEventRequest struct {
	Name          *string    `json:"name"`
	Description   *string    `json:"description"`
	Date          *time.Time `json:"date"`
	Location      *string    `json:"location"`
	InstitutionId *int64     `json:"institution_id"`
}

func (r UpdateEventRequest) ToModel() entities.Event {
	event := entities.Event{}
	if r.Name != nil {
		event.Name = *r.Name
	}
	if r.Description != nil {
		event.Description = *r.Description
	}
	if r.Date != nil {
		event.Date = *r.Date
	}
	if r.Location != nil {
		event.Location = *r.Location
	}
	if r.InstitutionId != nil {
		event.InstitutionId = *r.InstitutionId
	}
	return event
}
