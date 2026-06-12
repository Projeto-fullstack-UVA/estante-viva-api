package services

import (
	userdto "github.com/Projeto-fullstack-UVA/estante-viva-api/internals/dtos/users"
	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/repositories"
	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/utils"
)

// Login verifies the credentials and returns the matching user, or ErrUserNotFound.
func Login(email, password string) (*userdto.LoginResponse, error) {
	user, err := repositories.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}
	if err := utils.CheckPassword(user.Password, password); err != nil {
		return nil, ErrUserNotFound
	}
	
	resp, _ := userdto.NewLoginResponse(user)
	return &resp, nil
}

func Register(req userdto.CreateUserRequest) error {
	user := req.ToModel()

	hashed, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashed

	affected, err := repositories.CreateUser(user)
	if err != nil {
		return err
	}
	if affected == 0 {
		return ErrUserCreateFailed
	}
	return nil
}

func ListUsers() ([]userdto.UserResponse, error) {
	users, err := repositories.GetUsers()
	if err != nil {
		return nil, err
	}
	return userdto.NewUserResponseList(users), nil
}

// FindUser returns the user with the given id, or ErrUserNotFound.
func FindUser(id int64) (*userdto.UserResponse, error) {
	user, err := repositories.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}
	resp := userdto.NewUserResponse(*user)
	return &resp, nil
}
