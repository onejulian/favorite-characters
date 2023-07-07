package usecase

import (
	"mbs-back/src/domain"
)

func TokenIsValid(email, tokenString string) bool {
	isValid, err := domain.TokenIsValid(email, tokenString, tokenRepo)
	if err != nil {
		return false
	}
	return isValid
}
