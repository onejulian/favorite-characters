package usecase

import (
	"favorite-characters/src/domain"
	"favorite-characters/src/infraestructure/repository"
)

type GetUserUseCase struct {
}

func (g *GetUserUseCase) Execute(email string) (*domain.User, error) {
	result, err := domain.GetUser(email, repository.UserRepo)
	if err != nil {
		return nil, err
	}
	return result, nil
}
