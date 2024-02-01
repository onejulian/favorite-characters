package usecase

import (
	"favorite-characters/src/domain"
	"favorite-characters/src/infraestructure/repository"
)

type DeleteUserUseCase struct {
}

func (c *DeleteUserUseCase) Execute(email string) error {
	err := domain.DeleteUser(email, repository.UserRepo)
	if err != nil {
		return err
	}

	err = domain.DeleteAllCharacters(email, repository.CharacterRepo)
	if err != nil {
		return err
	}

	err = domain.DeleteAllTokens(email, repository.TokenRepo)
	if err != nil {
		return err
	}

	return nil
}
