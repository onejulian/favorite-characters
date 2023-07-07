package usecase

import (
	"mbs-back/src/domain"
)

type UpdateUserUseCase struct {
}

func (c *UpdateUserUseCase) Execute(user domain.User, token string) (*domain.User, error) {
	result, err := domain.UpdateUser(user, userRepo)
	if err != nil {
		return nil, err
	}
	return result, nil
}
