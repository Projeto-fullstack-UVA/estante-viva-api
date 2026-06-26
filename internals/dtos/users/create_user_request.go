package users

import (
	"time"

	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/entities"
)

type CreateUserRequest struct {
	Name          string    `json:"name" binding:"required"`
	Email         string    `json:"email" binding:"required,email"`
	BirthDate     time.Time `json:"birthDate" binding:"required"`
	Password      string    `json:"password" binding:"required"`
	Address       string    `json:"address" binding:"required"`
	Document      string    `json:"document" binding:"required,min=11,max=11,number"`
	Cellphone     string    `json:"cellphone" binding:"required,min=11,max=11,number"`
	Role          string    `json:"role" binding:"required,oneof=student teacher"`
	InstitutionID *int64    `json:"institution_id"`
	Score         int16     `json:"score"`
}

func (r CreateUserRequest) ToModel() entities.User {
	return entities.User{
		Name:          r.Name,
		Email:         r.Email,
		BirthDate:     r.BirthDate,
		Password:      r.Password,
		Address:       r.Address,
		Document:      r.Document,
		Cellphone:     r.Cellphone,
		Role:          r.Role,
		InstitutionID: r.InstitutionID,
		Score:         r.Score,
	}
}
