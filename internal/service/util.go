package service

import (
	"github.com/Pasca11/gRPC-Auth/models"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const secret = "secret" //TODO add env

func createToken(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": user.ID,
			"role":    user.Role,
			"exp":     time.Now().Add(time.Hour).Unix(),
		})
	signed, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return signed, nil
}
