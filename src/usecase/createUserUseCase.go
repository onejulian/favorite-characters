package usecase

import (
	"favorite-characters/src/domain"
)

type CreateUserUseCase struct {
}

func (c *CreateUserUseCase) Execute(user domain.User) (*domain.User, error) {
	result, err := domain.CreateUser(user, userRepo)
	if err != nil {
		return nil, err
	}
	return result, nil
}
