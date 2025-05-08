package services

import (
	"time"

	"github.com/AlyssonT/CheckpointBackend/communication/exceptions"
	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "secret"

type Jwt struct{}

func NewJwt() Jwt {
	return Jwt{}
}

func (j *Jwt) GenerateToken(email string, id uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"id":    id,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func (j *Jwt) VerifyToken(token string) (jwt.MapClaims, error) {
	parsedToken, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(t *jwt.Token) (any, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, exceptions.ErrorInvalidCredentials
		}

		return []byte(secretKey), nil
	})

	if err != nil || !parsedToken.Valid {
		return nil, exceptions.ErrorInvalidCredentials
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return nil, exceptions.ErrorInvalidCredentials
	}

	return claims, nil
}
