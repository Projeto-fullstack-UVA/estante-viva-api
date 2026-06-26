package institutions

import "github.com/Projeto-fullstack-UVA/estante-viva-api/internals/entities"

type UpdateInstitutionRequest struct {
	Name         *string `json:"name"`
	Abbreviation *string `json:"abbreviation"`
	City         *string `json:"city"`
	Address      *string `json:"address"`
}

func (r UpdateInstitutionRequest) ToModel() entities.Institution {
	institution := entities.Institution{}
	if r.Name != nil {
		institution.Name = *r.Name
	}
	if r.Abbreviation != nil {
		institution.Abbreviation = *r.Abbreviation
	}
	if r.City != nil {
		institution.City = *r.City
	}
	if r.Address != nil {
		institution.Address = *r.Address
	}
	return institution
}
