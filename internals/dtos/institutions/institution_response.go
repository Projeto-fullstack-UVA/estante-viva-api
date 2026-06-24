package institutions

import (
	"time"

	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/entities"
)

type InstitutionResponse struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	Abbreviation string    `json:"abbreviation"`
	City         string    `json:"city"`
	Address      string    `json:"address"`
	CreatedAt    time.Time `json:"created_at"`
}

func NewInstitutionResponse(i entities.Institution) InstitutionResponse {
	return InstitutionResponse{
		ID:           i.ID,
		Name:         i.Name,
		Abbreviation: i.Abbreviation,
		City:         i.City,
		Address:      i.Address,
		CreatedAt:    i.CreatedAt,
	}
}

func NewInstitutionResponseList(list []entities.Institution) []InstitutionResponse {
	out := make([]InstitutionResponse, 0, len(list))
	for _, i := range list {
		out = append(out, NewInstitutionResponse(i))
	}
	return out
}
