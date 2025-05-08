package services

import (
	"testing"

	"github.com/AlyssonT/CheckpointBackend/communication/exceptions"
	testutilities "github.com/AlyssonT/CheckpointBackend/test_utilities"
	"github.com/stretchr/testify/assert"
)

func TestJwt_Success(t *testing.T) {
	jwtService := NewJwt()
	user := testutilities.BuildFakeUser()
	id := 1

	token, err := jwtService.GenerateToken(user.Email, uint(id))

	assert.Nil(t, err)

	_, err = jwtService.VerifyToken(token)

	assert.Nil(t, err)
}

func TestJwtVerify_Fail(t *testing.T) {
	jwtService := NewJwt()

	_, err := jwtService.VerifyToken("invalidToken")

	assert.Equal(t, exceptions.ErrorInvalidCredentials, err)
}
