package usecase

import (
	"favorite-characters/src/domain"
	"favorite-characters/src/infraestructure/repository"
	"net/http"
)

type LogoutUseCase struct{}

func (l *LogoutUseCase) Execute(token domain.Token) (int, error) {
	err := domain.DeleteToken(token, repository.TokenRepo)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
