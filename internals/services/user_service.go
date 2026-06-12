package services

import (
	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/models"
	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/repositories"
)

// Login returns the user matching the given credentials, or ErrUserNotFound.
func Login(email, password string) (*models.User, error) {
	user, err := repositories.GetUserByCredentials(email, password)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}
	return user, nil
}

// Register creates a new user, returning ErrUserCreateFailed when nothing was inserted.
func Register(user models.User) error {
	affected, err := repositories.CreateUser(user)
	if err != nil {
		return err
	}
	if affected == 0 {
		return ErrUserCreateFailed
	}
	return nil
}

func ListUsers() ([]models.User, error) {
	return repositories.GetUsers()
}

// FindUser returns the user with the given id, or ErrUserNotFound.
func FindUser(id int64) (*models.User, error) {
	user, err := repositories.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrUserNotFound
	}
	return user, nil
}
