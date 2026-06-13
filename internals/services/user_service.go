package services

import (
	"log"

	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/auth"
	userdto "github.com/Projeto-fullstack-UVA/estante-viva-api/internals/dtos/users"
	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/repositories"
	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/utils"
)

func Login(email, password string) (userdto.LoginResponse, error) {
	user, err := repositories.GetUserByEmail(email)
	if err != nil {
		log.Println("Error while fetching user in the database: ", err.Error())
		return userdto.LoginResponse{}, err
	}
	if user == nil {
		log.Println("User not found in the database")
		return userdto.LoginResponse{}, ErrUserNotFound
	}
	if err := utils.CheckPassword(user.Password, password); err != nil {
		log.Println("The hashed password provided does not match with the one in the database")
		return userdto.LoginResponse{}, ErrUserNotFound
	}
	
	resp, err := userdto.NewLoginResponse(user)
	if err != nil {
		return userdto.LoginResponse{}, err
	}

	log.Println("Success logging user in")

	return resp, nil
}

func Register(req userdto.CreateUserRequest) (userdto.RegisterUserResponse, error) {
	user := req.ToModel()

	hashed, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Println("Error while hashing password: ", err)
		return userdto.RegisterUserResponse{}, ErrUserCreateFailed
	}
	user.Password = hashed

	affected, err := repositories.CreateUser(user)
	if err != nil {
		log.Println("Failed to create user in the database: ", err)
		return userdto.RegisterUserResponse{}, ErrUserCreateFailed
	}
	if affected == 0 {
		log.Println("Failed to create user in the database")
		return userdto.RegisterUserResponse{}, ErrUserCreateFailed
	}
	token, err := auth.GenerateToken(&user.ID, &user.Role)
	if err != nil {
		return userdto.RegisterUserResponse{}, ErrUserCreateFailed
	}

	resp := userdto.NewRegisterUserResponse(user)
	resp.Token = token

	log.Println("Success creating user in the database")
	
	return resp, nil
}

func ListUsers() ([]userdto.UserResponse, error) {
	users, err := repositories.GetUsers()
	if err != nil {
		log.Println("Error while fetching users from the database: ", err.Error())
		return nil, err
	}

	log.Println("Success fetching users")

	return userdto.NewListUserResponse(users), nil
}

func FindUser(id int64) (*userdto.UserResponse, error) {
	user, err := repositories.GetUserByID(id)
	if err != nil {
		log.Println("Error fetching user from the database: ", err.Error())
		return nil, err
	}
	if user == nil {
		log.Println("User was not found in the database")
		return nil, ErrUserNotFound
	}

	resp := userdto.NewUserResponse(*user)

	log.Println("Success fetching user")

	return &resp, nil
}
