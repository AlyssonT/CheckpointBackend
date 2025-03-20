package interfaces

type Cryptographer interface {
	HashPassword(password string) (string, error)
	CheckPassword(password string, hash string) bool
}
