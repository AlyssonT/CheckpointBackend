package services

import "golang.org/x/crypto/bcrypt"

type Cryptography struct {
	Cost int
}

const DefaultCost = 14

func NewCryptography(cost int) Cryptography {
	return Cryptography{
		Cost: cost,
	}
}

func (c *Cryptography) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), c.Cost)

	return string(bytes), err
}

func (c *Cryptography) CheckPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
