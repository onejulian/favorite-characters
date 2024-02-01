package usecase

import (
	"errors"
	"favorite-characters/src/domain"
	"favorite-characters/src/infraestructure/jwt"
	"favorite-characters/src/infraestructure/repository"
	"net/http"
)

type ChangePasswordUseCase struct {
}

func (c *ChangePasswordUseCase) Execute(email string, newPassword, oldPassword string) (int, error) {
	user, err := domain.GetUser(email, repository.UserRepo)
	if err != nil {
		return http.StatusNotFound, err
	}

	if !jwt.CheckPasswordHash(user.Password, []byte(oldPassword)) {
		return http.StatusUnauthorized, errors.New("la contraseña anterior es incorrecta")
	}

	err = domain.ChangePassword(email, newPassword, repository.UserRepo)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	err = domain.DeleteAllTokens(email, repository.TokenRepo)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
