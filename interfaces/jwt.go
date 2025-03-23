package interfaces

type JwtService interface {
	GenerateToken(email string, id uint) (string, error)
	VerifyToken(token string) error
}
