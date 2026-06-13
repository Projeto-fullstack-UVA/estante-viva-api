package services

import (
	"errors"
	"log"

	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/auth"
	userdto "github.com/Projeto-fullstack-UVA/estante-viva-api/internals/dtos/users"
	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/repositories"
	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/utils"
)

func Login(email, password string) (userdto.LoginResponse, error) {
	user, err := repositories.GetUserByEmail(email)
	if err != nil {
		return userdto.LoginResponse{}, err
	}
	if user == nil {
		return userdto.LoginResponse{}, ErrUserNotFound
	}
	if err := utils.CheckPassword(user.Password, password); err != nil {
		return userdto.LoginResponse{}, ErrUserNotFound
	}
	
	resp, err := userdto.NewLoginResponse(user)
	if err != nil {
		return userdto.LoginResponse{}, err
	}
	return resp, nil
}

func Register(req userdto.CreateUserRequest) (userdto.RegisterUserResponse, error) {
	user := req.ToModel()

	hashed, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Println("unsupported password hash: ", err)
		return userdto.RegisterUserResponse{}, errors.New("unsupported password hash")
	}
	user.Password = hashed

	affected, err := repositories.CreateUser(user)
	if err != nil {
		log.Println("failed to register user in the database: ", err)
		return userdto.RegisterUserResponse{}, errors.New("failed to register user in the database")
	}
	if affected == 0 {
		return userdto.RegisterUserResponse{}, ErrUserCreateFailed
	}
	token, err := auth.GenerateToken(&user.ID, &user.Role)
	if err != nil {
		return userdto.RegisterUserResponse{}, err
	}
	resp := userdto.NewRegisterUserResponse(user)
	resp.Token = token
	return resp, nil
}

func ListUsers() ([]userdto.UserResponse, error) {
	users, err := repositories.GetUsers()
	if err != nil {
		return nil, err
	}
	return userdto.NewListUserResponse(users), nil
}

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
