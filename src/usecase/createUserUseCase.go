package usecase

import (
	"mbs-back/src/domain"
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
