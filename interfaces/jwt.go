package interfaces

import "github.com/golang-jwt/jwt/v5"

type JwtService interface {
	GenerateToken(name string, email string, id uint) (string, error)
	VerifyToken(token string) (jwt.MapClaims, error)
}
