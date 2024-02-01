package repository

import (
	dynamodbDao "favorite-characters/src/dao/dynamodbDao"
	postgresDao "favorite-characters/src/dao/postgresqlDao"
	"favorite-characters/src/domain"
	"os"
)

var UserRepo domain.UserRepository
var CharacterRepo domain.CharacterRepository
var TokenRepo domain.TokenRepository

func addStage() string {
	env := os.Getenv("ENV")
	if env == "DEV" {
		return "_stage"
	}
	return ""
}

func init() {
	dbEngine := os.Getenv("DB_ENGINE")
	switch dbEngine {
	case "postgres":
		UserRepo = postgresDao.NewUserDao()
		CharacterRepo = postgresDao.NewCharacterDao()
		TokenRepo = postgresDao.NewTokenDao()
	default:
		stage := addStage()
		UserRepo = dynamodbDao.NewUserDao("users" + stage)
		CharacterRepo = dynamodbDao.NewCharacterDao("user_favorites" + stage)
		TokenRepo = dynamodbDao.NewTokenDao("tokens" + stage)
	}
}
