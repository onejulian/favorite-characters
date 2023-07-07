package usecase

import "mbs-back/src/domain"

type AddCharacterUseCase struct {
}

func (a *AddCharacterUseCase) Execute(userEmail string, idCharacter string) (*domain.Character, error) {
	return domain.CreateCharacter(domain.Character{
		UserEmail:   userEmail,
		IdCharacter: idCharacter,
	}, characterRepo)
}
