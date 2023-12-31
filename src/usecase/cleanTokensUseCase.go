package usecase

import (
	"mbs-back/src/domain"
	"mbs-back/src/infraestructure/jwt"
)

func cleanTokens(email string) error {
	allTokens, err := domain.GetToken(email, tokenRepo)
	if err != nil {
		return err
	}

	for _, token := range *allTokens {
		isValid, _ := jwt.ValidateJWT(token.Value)

		if !isValid {
			err = domain.DeleteToken(token, tokenRepo)
			if err != nil {
				return err
			}
		}

	}

	return nil
}
