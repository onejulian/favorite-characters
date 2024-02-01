package usecase

import (
	"favorite-characters/src/domain"
	"favorite-characters/src/infraestructure/repository"
)

type CreateUserUseCase struct {
}

func (c *CreateUserUseCase) Execute(user domain.User) (*domain.User, error) {
	result, err := domain.CreateUser(user, repository.UserRepo)
	if err != nil {
		return nil, err
	}
	return result, nil
}
