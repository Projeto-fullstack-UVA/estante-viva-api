package institutions

import "github.com/Projeto-fullstack-UVA/estante-viva-api/internals/entities"

type CreateInstitutionRequest struct {
	Name         string `json:"name" binding:"required"`
	Abbreviation string `json:"abbreviation" binding:"required"`
	City         string `json:"city" binding:"required"`
	Address      string `json:"address" binding:"required"`
}

func (i CreateInstitutionRequest) ToModel() entities.Institution {
	return entities.Institution {
		Name: i.Name,
		Abbreviation: i.Abbreviation,
		City: i.City,
		Address: i.Address,
	}
}