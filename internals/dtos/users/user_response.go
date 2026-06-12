package users

import (
	"time"

	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/models"
)

// UserResponse is the user representation returned to clients (never the password).
type UserResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Address   string    `json:"address"`
	Document  string    `json:"document"`
	Cellphone string    `json:"cellphone"`
	Role      string    `json:"role"`
	Campus    string    `json:"campus"`
	Score     int16     `json:"score"`
	CreatedAt time.Time `json:"created_at"`
}

func NewUserResponse(u models.User) UserResponse {
	return UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Address:   u.Address,
		Document:  u.Document,
		Cellphone: u.Cellphone,
		Role:      u.Role,
		Campus:    u.Campus,
		Score:     u.Score,
		CreatedAt: u.CreatedAt,
	}
}

func NewUserResponseList(list []models.User) []UserResponse {
	out := make([]UserResponse, 0, len(list))
	for _, u := range list {
		out = append(out, NewUserResponse(u))
	}
	return out
}
