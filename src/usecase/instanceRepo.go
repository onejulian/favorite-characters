package usecase

import (
	"mbs-back/src/dao"
	"mbs-back/src/domain"
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
