package services

import (
	"time"

	"github.com/AlyssonT/CheckpointBackend/communication/exceptions"
	"github.com/AlyssonT/CheckpointBackend/interfaces"
	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "secret"

type Jwt struct{}

func NewJwt() Jwt {
	return Jwt{}
}

func (j *Jwt) ExtractClaims(claims map[string]any) (*interfaces.UserClaims, error) {
	name, ok := claims["name"].(string)
	if !ok {
		return nil, exceptions.ErrorInvalidCredentials
	}

	email, ok := claims["email"].(string)
	if !ok {
		return nil, exceptions.ErrorInvalidCredentials
	}

	idFloat, ok := claims["id"].(float64)
	if !ok {
		return nil, exceptions.ErrorInvalidCredentials
	}
	id := uint(idFloat)

	expFloat, ok := claims["exp"].(float64)
	if !ok {
		return nil, exceptions.ErrorInvalidCredentials
	}
	exp := int64(expFloat)

	return &interfaces.UserClaims{
		Name:  name,
		Email: email,
		ID:    id,
		Exp:   exp,
	}, nil
}

func (j *Jwt) GenerateToken(name string, email string, id uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":  name,
		"email": email,
		"id":    id,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func (j *Jwt) VerifyToken(token string) (*interfaces.UserClaims, error) {
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

	mappedClaims, err := j.ExtractClaims(claims)

	if err != nil {
		return nil, exceptions.ErrorInvalidCredentials
	}

	return mappedClaims, nil
}
