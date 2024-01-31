package dynamo_test

import (
	"favorite-characters/src/domain"
	"favorite-characters/src/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	assert := assert.New(t)

	user := domain.User{
		Email:     "usertest@mail.com",
		FirtsName: "user",
		LastName:  "test",
		Password:  "123456",
		IsActive:  true,
	}

	createUserUsecase := usecase.CreateUserUseCase{}
	_, err := createUserUsecase.Execute(user)

	assert.Nil(err)
}

func TestDeleteUser(t *testing.T) {
	assert := assert.New(t)

	email := "usertest@mail.com"

	deleteUserUsecase := usecase.DeleteUserUseCase{}
	err := deleteUserUsecase.Execute(email)

	assert.Nil(err)
}
