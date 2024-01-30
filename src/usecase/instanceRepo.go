package usecase

import (
	"favorite-characters/src/dao"
	"favorite-characters/src/domain"
	"os"
)

var stage = addStage()
var userRepo domain.UserRepository = dao.NewUserDao("users" + stage)
var characterRepo domain.CharacterRepository = dao.NewCharacterDao("user_favorites" + stage)
var tokenRepo domain.TokenRepository = dao.NewTokenDao("tokens" + stage)

func addStage() string {
	env := os.Getenv("ENV")
	if env == "DEV" {
		return "_stage"
	}
	return ""
}
