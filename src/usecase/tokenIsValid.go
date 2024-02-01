package usecase

import (
	"favorite-characters/src/domain"
	"favorite-characters/src/infraestructure/repository"
)

func TokenIsValid(email, tokenString string) bool {
	isValid, err := domain.TokenIsValid(email, tokenString, repository.TokenRepo)
	if err != nil {
		return false
	}
	return isValid
}
