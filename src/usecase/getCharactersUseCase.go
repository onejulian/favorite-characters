package usecase

import "mbs-back/src/domain"

type GetCharactersUseCase struct {
}

func (a *GetCharactersUseCase) Execute(userEmail string) ([]*domain.Character, error) {
	go cleanTokens(userEmail)
	return domain.GetCharacters(userEmail, characterRepo)
}
