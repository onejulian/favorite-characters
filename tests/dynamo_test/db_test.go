package db_test

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

func TestLoginLogout(t *testing.T) {
	assert := assert.New(t)

	email := "usertest@mail.com"
	password := "123456"

	loginUsecase := usecase.LoginUseCase{}
	resp, err := loginUsecase.Execute(email, password)

	assert.Nil(err)
	assert.NotNil(resp)

	logoutUsecase := usecase.LogoutUseCase{}
	token := domain.Token{Value: resp.Token, UserEmail: email}
	code, err := logoutUsecase.Execute(token)

	assert.Nil(err)
	assert.Equal(200, code)
}

func TestDeleteUser(t *testing.T) {
	assert := assert.New(t)

	email := "usertest@mail.com"

	deleteUserUsecase := usecase.DeleteUserUseCase{}
	err := deleteUserUsecase.Execute(email)

	assert.Nil(err)
}
