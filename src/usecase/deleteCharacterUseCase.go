package usecase

import (
	"favorite-characters/src/domain"
	"favorite-characters/src/infraestructure/repository"
)

type DeleteCharacterUseCase struct {
}

func (a *DeleteCharacterUseCase) Execute(userEmail string, idCharacter string) error {
	return domain.DeleteCharacter(userEmail, idCharacter, repository.CharacterRepo)
}
