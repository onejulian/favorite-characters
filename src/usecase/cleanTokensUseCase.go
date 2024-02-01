package usecase

import (
	"favorite-characters/src/domain"
	"favorite-characters/src/infraestructure/jwt"
	"favorite-characters/src/infraestructure/repository"
)

func cleanTokens(email string) error {
	allTokens, err := domain.GetToken(email, repository.TokenRepo)
	if err != nil {
		return err
	}

	for _, token := range *allTokens {
		isValid, _ := jwt.ValidateJWT(token.Value)

		if !isValid {
			err = domain.DeleteToken(token, repository.TokenRepo)
			if err != nil {
				return err
			}
		}

	}

	return nil
}
