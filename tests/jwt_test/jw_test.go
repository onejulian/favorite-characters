package jwt_test

import (
	"favorite-characters/src/infraestructure/jwt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateJWT(t *testing.T) {
	assert := assert.New(t)

	password_encoded := jwt.HashAndSalt([]byte("holamundo"))
	assert.Empty(password_encoded)
}