package services

import (
	"context"
	"errors"
	"log"

	institutionDto "github.com/Projeto-fullstack-UVA/estante-viva-api/internals/dtos/institutions"
	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/repositories"
)

func ListInstitutions(ctx context.Context) ([]institutionDto.InstitutionResponse, error) {
	institutions, err := repositories.GetInstitutions(ctx)
	if err != nil {
		log.Println("Error while fetching institutions from the database: ", err.Error())
		return nil, errors.New("Failed to get institutions")
	}

	log.Println("Success fetching institutions from the database")
	return institutionDto.NewInstitutionResponseList(institutions), nil
}

func FindInstitution(ctx context.Context, id int64) (*institutionDto.InstitutionResponse, error) {
	institution, err := repositories.GetInstitutionById(ctx, id)
	if err != nil {
		log.Println("Error while fetching institution by id: ", err.Error())
		return nil, ErrInstitutionFetchFailed
	}
	if institution == nil {
		log.Println("No institution was found with the provided id")
		return nil, ErrInstitutionNotFound
	}

	log.Println("Success fetching institution from database")

	resp := institutionDto.NewInstitutionResponse(*institution)
	return &resp, nil
}
