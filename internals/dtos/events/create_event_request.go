package events

import (
	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/entities"
)

type CreateEventRequest struct {
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Location      string    `json:"location"`
	InstitutionId int64     `json:"institution_id"`
}

func (r CreateEventRequest) ToModel() entities.Event {
	return entities.Event{
		Name: r.Name,
		Description: r.Description,
		Location: r.Location,
		InstitutionId: r.InstitutionId,
	}
}