package controllers

import (
	"errors"
	"net/http"
	"strconv"

	userdto "github.com/Projeto-fullstack-UVA/estante-viva-api/internals/dtos/users"
	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/middleware"
	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/services"
	"github.com/Projeto-fullstack-UVA/estante-viva-api/internals/utils"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var req userdto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Fail(c, http.StatusBadRequest, "INVALID_CREDENTIALS", "Invalid credentials format")
		return
	}

	user, err := services.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			utils.Fail(c, http.StatusUnauthorized, "INVALID_CREDENTIALS", "Invalid email or password")
			return
		}
		utils.Fail(c, http.StatusInternalServerError, "INTERNAL_ERROR", "Error while logging in")
		return
	}

	c.Header("Authorization", user.Token)
	utils.OK(c, http.StatusOK, user)
}

func Register(c *gin.Context) {
	var req userdto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Fail(c, http.StatusBadRequest, "INVALID_PAYLOAD", "Invalid user format")
		return
	}

	resp, err := services.Register(c.Request.Context(), req)
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, "INTERNAL_ERROR", "Error while creating user")
		return
	}

	utils.OK(c, http.StatusCreated, resp)
}

func GetMe(c *gin.Context) {
	id, ok := middleware.GetCurrentUserID(c)
	if !ok {
		utils.Fail(c, http.StatusUnauthorized, "UNAUTHORIZED", "Unauthorized")
		return
	}

	user, err := services.FindUser(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			utils.Fail(c, http.StatusNotFound, "USER_NOT_FOUND", "User not found")
			return
		}
		utils.Fail(c, http.StatusInternalServerError, "INTERNAL_ERROR", "Error while finding user")
		return
	}

	utils.OK(c, http.StatusOK, user)
}

func ListUsers(c *gin.Context) {
	users, err := services.ListUsers(c.Request.Context())
	if err != nil {
		utils.Fail(c, http.StatusInternalServerError, "INTERNAL_ERROR", "Error while returning the user's list")
		return
	}

	utils.OK(c, http.StatusOK, users)
}

func FindUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.Fail(c, http.StatusBadRequest, "INVALID_ID", "Invalid ID")
		return
	}

	user, err := services.FindUser(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			utils.Fail(c, http.StatusNotFound, "USER_NOT_FOUND", "User not found")
			return
		}
		utils.Fail(c, http.StatusInternalServerError, "INTERNAL_ERROR", "Error while finding user")
		return
	}

	utils.OK(c, http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.Fail(c, http.StatusBadRequest, "INVALID_ID", "Invalid ID")
		return
	}

	var req userdto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Fail(c, http.StatusBadRequest, "INVALID_PAYLOAD", "Invalid user format")
		return
	}

	if err := services.UpdateUser(c.Request.Context(), id, req); err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			utils.Fail(c, http.StatusNotFound, "USER_NOT_FOUND", "User not found")
			return
		}
		utils.Fail(c, http.StatusInternalServerError, "INTERNAL_ERROR", "Error while updating user")
		return
	}

	utils.OK(c, http.StatusOK, gin.H{"message": "User updated successfully"})
}

func DeleteUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.Fail(c, http.StatusBadRequest, "INVALID_ID", "Invalid ID")
		return
	}

	if err := services.DeleteUser(c.Request.Context(), id); err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			utils.Fail(c, http.StatusNotFound, "USER_NOT_FOUND", "User not found")
			return
		}
		utils.Fail(c, http.StatusInternalServerError, "INTERNAL_ERROR", "Error while deleting user")
		return
	}

	utils.OK(c, http.StatusOK, gin.H{"message": "User deleted successfully"})
}
