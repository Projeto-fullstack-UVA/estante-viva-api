package events

import (
	"time"

	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/entities"
)

type EventResponse struct {
	ID            int64      `json:"id"`
	Name          string     `json:"name"`
	Description   string     `json:"description"`
	Date          time.Time  `json:"date"`
	Location      string     `json:"location"`
	InstitutionId int64      `json:"intitution_id"`
	CreatedAt     *time.Time `json:"created_at"`
}

func NewEventResponse(e entities.Event) EventResponse {
	return EventResponse{
		ID:            e.ID,
		Name:          e.Name,
		Description:   e.Description,
		Date:          e.Date,
		Location:      e.Location,
		InstitutionId: e.InstitutionId,
		CreatedAt:     e.CreatedAt,
	}
}

func NewEventResponseList(list []entities.Event) []EventResponse {
	out := make([]EventResponse, 0, len(list))
	for _, e := range list {
		out = append(out, NewEventResponse(e))
	}
	return out
}
