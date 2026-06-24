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
			c.String(http.StatusUnauthorized, "Invalid email or password")
			return
		}
		c.String(http.StatusInternalServerError, "Error while logging in")
		return
	}

	c.Header("Authorization", user.Token)
	c.JSON(http.StatusOK, user)
}

func Register(c *gin.Context) {
	var req userdto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.String(http.StatusBadRequest, "Invalid user format")
		return
	}

	resp, err := services.Register(req)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error while creating user")
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func ListUsers(c *gin.Context) {
	users, err := services.ListUsers()
	if err != nil {
		c.String(http.StatusInternalServerError, "Error while returning the user's list")
		return
	}

	c.JSON(http.StatusOK, users)
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
		c.String(http.StatusInternalServerError, "Error while finding user")
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid ID")
		return
	}

	var req userdto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.String(http.StatusBadRequest, "Invalid user format")
		return
	}

	if err := services.UpdateUser(id, req); err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			c.String(http.StatusNotFound, "User not found")
			return
		}
		c.String(http.StatusInternalServerError, "Error while updating user")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "success": true})
}

func DeleteUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid ID")
		return
	}

	if err := services.DeleteUser(id); err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			c.String(http.StatusNotFound, "User not found")
			return
		}
		c.String(http.StatusInternalServerError, "Error while deleting user")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully", "success": true})
}
