package repository

import (
	dynamodbDao "favorite-characters/src/dao/dynamodbDao"
	mysqlDao "favorite-characters/src/dao/mysqlDao"
	"favorite-characters/src/domain"
	"favorite-characters/src/infraestructure/dbconfig/dynamoconfig"
	"favorite-characters/src/infraestructure/dbconfig/mysqlconfig"
	"favorite-characters/src/infraestructure/env"
	"os"
)

var UserRepo domain.UserRepository
var CharacterRepo domain.CharacterRepository
var TokenRepo domain.TokenRepository

func getStage() string {
	env := os.Getenv("ENV")
	if env == "DEV" {
		return "_stage"
	}
	return ""
}

func init() {
	env.LoadEnv()
	dbEngine := os.Getenv("DB_ENGINE")
	switch dbEngine {
	case "mysql":
		mysqlconfig.InitMysqlDB()
		UserRepo = mysqlDao.NewUserDao()
		CharacterRepo = mysqlDao.NewCharacterDao()
		TokenRepo = mysqlDao.NewTokenDao()
	default:
		stage := getStage()
		dynamoconfig.InitDynaClient()
		UserRepo = dynamodbDao.NewUserDao("users" + stage)
		CharacterRepo = dynamodbDao.NewCharacterDao("user_favorites" + stage)
		TokenRepo = dynamodbDao.NewTokenDao("tokens" + stage)
	}
}
