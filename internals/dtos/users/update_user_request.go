package users

import (
	"time"

	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/entities"
)

type UpdateUserRequest struct {
	Name          *string    `json:"name"`
	Email         *string    `json:"email"`
	Address       *string    `json:"address"`
	Document      *string    `json:"document"`
	Cellphone     *string    `json:"cellphone"`
	InstitutionID *int64     `json:"institution_id"`
	BirthDate     *time.Time `json:"birthDate"`
}

func (r UpdateUserRequest) ToModel() entities.User {
	user := entities.User{}
	if r.Name != nil {
		user.Name = *r.Name
	}
	if r.Email != nil {
		user.Email = *r.Email
	}
	if r.Address != nil {
		user.Address = *r.Address
	}
	if r.Document != nil {
		user.Document = *r.Document
	}
	if r.Cellphone != nil {
		user.Cellphone = *r.Cellphone
	}
	if r.InstitutionID != nil {
		user.InstitutionID = r.InstitutionID
	}
	if r.BirthDate != nil {
		user.BirthDate = *r.BirthDate
	}
	return user
}
