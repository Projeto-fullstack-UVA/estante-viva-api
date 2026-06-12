package controllers

import (
	"errors"
	"net/http"
	"strconv"

	userdto "github.com/Projeto-fullstack-UVA/estante-viva-api/internals/dtos/users"
	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/services"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var req userdto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.String(http.StatusBadRequest, "Invalid credentials format")
		return
	}

	user, err := services.Login(req.Email, req.Password)
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			c.String(http.StatusNotFound, "User not found")
			return
		}
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, userdto.NewUserResponse(*user))
}

func Register(c *gin.Context) {
	var req userdto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.String(http.StatusBadRequest, "Invalid user format")
		return
	}

	if err := services.Register(req.ToModel()); err != nil {
		c.String(http.StatusInternalServerError, "Error while creating user")
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "success": true})
}

func ListUsers(c *gin.Context) {
	users, err := services.ListUsers()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, userdto.NewUserResponseList(users))
}

func FindUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid ID")
		return
	}

	user, err := services.FindUser(id)
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			c.String(http.StatusNotFound, "User not found")
			return
		}
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, userdto.NewUserResponse(*user))
}
