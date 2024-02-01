package usecase

import (
	"favorite-characters/src/domain"
	"favorite-characters/src/infraestructure/repository"
)

type UpdateUserUseCase struct {
}

func (c *UpdateUserUseCase) Execute(user domain.User, token string) (*domain.User, error) {
	result, err := domain.UpdateUser(user, repository.UserRepo)
	if err != nil {
		return nil, err
	}
	return result, nil
}
