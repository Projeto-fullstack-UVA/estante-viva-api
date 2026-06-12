package users

import (
	"time"

	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/models"
)

// CreateUserRequest is the expected body for POST /users.
type CreateUserRequest struct {
	Name      string    `json:"name" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	Address   string    `json:"address"`
	Document  string    `json:"document"`
	Cellphone string    `json:"cellphone"`
	Role      string    `json:"role" binding:"required,oneof=student teacher donator admin"`
	Campus    string    `json:"campus"`
	Score     int16     `json:"score"`
	CreatedAt time.Time `json:"created_at" binding:"required"`
}

func (r CreateUserRequest) ToModel() models.User {
	return models.User{
		Name:      r.Name,
		Email:     r.Email,
		Password:  r.Password,
		Address:   r.Address,
		Document:  r.Document,
		Cellphone: r.Cellphone,
		Role:      r.Role,
		Campus:    r.Campus,
		Score:     r.Score,
		CreatedAt: r.CreatedAt,
	}
}
