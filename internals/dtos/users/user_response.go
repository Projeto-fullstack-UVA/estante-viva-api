package users

import (
	"errors"
	"time"

	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/auth"
	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/entities"
)

type RegisterUserResponse struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Token     string    `json:"token"`
}

type LoginResponse struct {
	ID    int64  `json:"id"`
	Token string `json:"token"`
}

type UserResponse struct {
	ID        int64     `json:"id"`
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

func NewRegisterUserResponse(u entities.User) RegisterUserResponse {
	return RegisterUserResponse{
		ID:        u.ID,
		CreatedAt: u.CreatedAt,
	}
}

func NewListUserResponse(list []entities.User) []UserResponse {
	out := make([]UserResponse, 0, len(list))
	for _, u := range list {
		out = append(out, UserResponse{
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
		})
	}
	return out
}

func NewLoginResponse(u *entities.User) (LoginResponse, error) {
	token, err := auth.GenerateToken(&u.ID, &u.Role)
	if err != nil {
		return LoginResponse{}, errors.New("Login failed")
	}

	return LoginResponse{
		ID:    u.ID,
		Token: token,
	}, nil
}

func NewUserResponse(u entities.User) UserResponse {
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

func NewRegisterUserResponseList(list []entities.User) []RegisterUserResponse {
	out := make([]RegisterUserResponse, 0, len(list))
	for _, u := range list {
		out = append(out, NewRegisterUserResponse(u))
	}
	return out
}
