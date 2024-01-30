package usecase

import (
	"errors"
	"favorite-characters/src/domain"
	"favorite-characters/src/infraestructure/jwt"
	"net/http"
)

type ChangePasswordUseCase struct {
}

func (c *ChangePasswordUseCase) Execute(email string, newPassword, oldPassword string) (int, error) {
	user, err := domain.GetUser(email, userRepo)
	if err != nil {
		return http.StatusNotFound, err
	}

	if !jwt.CheckPasswordHash(user.Password, []byte(oldPassword)) {
		return http.StatusUnauthorized, errors.New("la contraseña anterior es incorrecta")
	}

	err = domain.ChangePassword(email, newPassword, userRepo)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	err = domain.DeleteAllTokens(email, tokenRepo)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
