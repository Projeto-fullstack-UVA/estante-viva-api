package auth

import (
	"strings"
	"time"

	"errors"

	environment "github.com/Projeto-fullstack-UVA/estante-viva-api/internals/config"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID *int64, role *string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(environment.JwtSecretKey))

	if err != nil {
		return "", errors.New("Failed to sign token: " + err.Error())
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	tokenString = RemoveBearerPrefix(tokenString)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}

		return []byte(environment.JwtSecretKey), nil
	})
	if err != nil {
		return nil, errors.New("Invalid token: " + err.Error())
	}

	return token, nil
}

func RemoveBearerPrefix(tokenString string) string {
	return strings.TrimPrefix(tokenString, "Bearer ")
}
