package usecase

import (
	"errors"
	"favorite-characters/src/domain"
	"favorite-characters/src/infraestructure/jwt"
	"favorite-characters/src/infraestructure/repository"
	"net/http"
)

type LoginUseCase struct{}

func (useCase *LoginUseCase) Execute(email, password string) (jwt.ResponseLogin, error) {
	user, err := domain.GetUser(email, repository.UserRepo)
	if err != nil {
		return jwt.ResponseLogin{
			Code: http.StatusNotFound,
		}, err
	}

	if !jwt.CheckPasswordHash(user.Password, []byte(password)) {
		return jwt.ResponseLogin{
			Code: http.StatusUnauthorized,
		}, errors.New("invalid password")
	}

	go cleanTokens(email)

	token, err := jwt.CreateJWT(email)
	if err != nil {
		return jwt.ResponseLogin{
			Code: http.StatusInternalServerError,
		}, err
	}

	tokenStr := domain.Token{
		UserEmail: email,
		Value:     token,
	}

	_, err = domain.CreateToken(tokenStr, repository.TokenRepo)
	if err != nil {
		return jwt.ResponseLogin{
			Code: http.StatusInternalServerError,
		}, err
	}

	return jwt.ResponseLogin{
		Email: email,
		Token: token,
		Code:  http.StatusOK,
	}, nil
}
