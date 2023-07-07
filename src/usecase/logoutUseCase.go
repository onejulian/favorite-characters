package usecase

import (
	"mbs-back/src/domain"
	"net/http"
)

type LogoutUseCase struct{}

func (l *LogoutUseCase) Execute(token domain.Token) (int, error) {
	err := domain.DeleteToken(token, tokenRepo)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
