package services

import (
	"testing"

	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/assert"
)

func TestCryptography(t *testing.T) {
	f := faker.New()
	password := f.Internet().Password()

	crypt := NewCryptography(DefaultCost)
	hashedPassword, err := crypt.HashPassword(password)

	assert.Nil(t, err)

	isValidated := crypt.CheckPassword(hashedPassword, password)

	assert.True(t, isValidated)
}
