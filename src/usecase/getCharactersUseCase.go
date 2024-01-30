package usecase

import "favorite-characters/src/domain"

type GetCharactersUseCase struct {
}

func (a *GetCharactersUseCase) Execute(userEmail string) ([]*domain.Character, error) {
	go cleanTokens(userEmail)
	return domain.GetCharacters(userEmail, characterRepo)
}
