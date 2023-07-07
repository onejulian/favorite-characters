package usecase

import "mbs-back/src/domain"

type DeleteCharacterUseCase struct {
}

func (a *DeleteCharacterUseCase) Execute(userEmail string, idCharacter string) error {
	return domain.DeleteCharacter(userEmail, idCharacter, characterRepo)
}
