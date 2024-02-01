package usecase

import (
	"favorite-characters/src/domain"
	"favorite-characters/src/infraestructure/repository"
)

type AddCharacterUseCase struct {
}

func (a *AddCharacterUseCase) Execute(userEmail string, idCharacter string) (*domain.Character, error) {
	return domain.CreateCharacter(domain.Character{
		UserEmail:   userEmail,
		IdCharacter: idCharacter,
	}, repository.CharacterRepo)
}
