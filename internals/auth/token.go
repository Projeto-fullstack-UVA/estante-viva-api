package auth

import (
	"os"
	"strings"
	"time"

	"errors"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID *int64, role *string) (string, error) {
	secret := os.Getenv("JWT_SECRET_KEY")
	if secret == "" {
		return "", errors.New("Variable JWT_SECRET_KEY is not set")
	}

	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", errors.New("Failed to sign token: " + err.Error())
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	tokenString = RemoveBearerPrefix(tokenString)

	secret := os.Getenv("JWT_SECRET_KEY")
	if secret == "" {
		return nil, errors.New("Variable JWT_SECRET_KEY is not set")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}

		return []byte(secret), nil
	})
	if err != nil {
		return nil, errors.New("Invalid token: " + err.Error())
	}

	return token, nil
}

func RemoveBearerPrefix(tokenString string) string {
	return strings.TrimPrefix(tokenString, "Bearer ")
}
