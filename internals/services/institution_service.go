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

func CreateInstitution(ctx context.Context, req institutionDto.CreateInstitutionRequest) error {
	affected, err := repositories.CreateInstitution(ctx, req.ToModel())
	if err != nil {
		log.Println("Error while creating institution in the database: ", err)
		return ErrInstitutionCreateFailed
	}
	if affected == 0 {
		log.Println("Failed to register institution")
		return ErrInstitutionCreateFailed
	}

	log.Println("Success creating institution in the database")
	return nil
}

func DeleteInstitution(ctx context.Context, id int64) error {
	affected, err := repositories.DeleteInstitution(ctx, id)
	if err != nil {
		log.Println("Error while deleting institution in the database: ", err)
		return ErrInstitutionDeleteFailed
	}
	if affected == 0 {
		log.Println("No institution found to delete with the provided id")
		return ErrInstitutionNotFound
	}

	log.Println("Institution deleted successfully")
	return nil
}
