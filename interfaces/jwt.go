package interfaces

type UserClaims struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	ID    uint   `json:"id"`
	Exp   int64  `json:"exp"`
}

type JwtService interface {
	GenerateToken(name string, email string, id uint) (string, error)
	VerifyToken(token string) (*UserClaims, error)
	ExtractClaims(claims map[string]any) (*UserClaims, error)
}
