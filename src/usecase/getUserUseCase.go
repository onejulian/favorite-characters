package usecase

import (
	"favorite-characters/src/domain"
)

type GetUserUseCase struct {
}

func (g *GetUserUseCase) Execute(email string) (*domain.User, error) {
	result, err := domain.GetUser(email, userRepo)
	if err != nil {
		return nil, err
	}
	return result, nil
}
