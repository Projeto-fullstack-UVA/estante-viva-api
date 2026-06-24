package services

import (
	"errors"
	"log"

	institutionDto "github.com/Projeto-fullstack-UVA/estante-viva-api/internals/dtos/institutions"
	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/repositories"
)

func ListInstitutions() ([]institutionDto.InstitutionResponse, error) {
	institutions, err := repositories.GetInstitutions()
	if err != nil {
		log.Println("Error while fetching institutions from the database: ", err.Error())
		return nil, errors.New("Failed to get institutions")
	}
	
	log.Println("Success fetching institutions from the database")
	return institutionDto.NewInstitutionResponseList(institutions), nil
}
