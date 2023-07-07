package usecase

import (
	"mbs-back/src/domain"
)

type DeleteUserUseCase struct {
}

func (c *DeleteUserUseCase) Execute(email string) error {
	err := domain.DeleteUser(email, userRepo)
	if err != nil {
		return err
	}

	err = domain.DeleteAllCharacters(email, characterRepo)
	if err != nil {
		return err
	}

	err = domain.DeleteAllTokens(email, tokenRepo)
	if err != nil {
		return err
	}

	return nil
}
