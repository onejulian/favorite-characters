package view

import (
	"favorite-characters/src/domain"
	"favorite-characters/src/infraestructure/jwt"
	"favorite-characters/src/usecase"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/gin-gonic/gin"
)

var (
	ErrorIdIsRequiered = "el id es requerido"
)

func AddCharacter(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": ErrorIdIsRequiered})
		return
	}

	email, err := jwt.EmailFromToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorBody{ErrorMsg: aws.String(err.Error())})
		return
	}
	token := jwt.GetTokenRequest(c)
	if !usecase.TokenIsValid(email, token) {
		c.JSON(http.StatusBadRequest, domain.ErrorBody{ErrorMsg: aws.String(ErrorTokenInvalid)})
		return
	}
	addCharacterUsecase := usecase.AddCharacterUseCase{}
	result, err := addCharacterUsecase.Execute(email, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorBody{ErrorMsg: aws.String(err.Error())})
		return
	}

	c.JSON(201, result)
}

func DeleteCharacter(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": ErrorIdIsRequiered})
		return
	}
	email, err := jwt.EmailFromToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorBody{ErrorMsg: aws.String(err.Error())})
		return
	}
	token := jwt.GetTokenRequest(c)
	if !usecase.TokenIsValid(email, token) {
		c.JSON(http.StatusBadRequest, domain.ErrorBody{ErrorMsg: aws.String(ErrorTokenInvalid)})
		return
	}

	deleteCharacterUsecase := usecase.DeleteCharacterUseCase{}
	err = deleteCharacterUsecase.Execute(email, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorBody{ErrorMsg: aws.String(err.Error())})
		return
	}

	c.JSON(200, gin.H{"message": "Personaje eliminado"})
}

func GetCharacters(c *gin.Context) {
	email, err := jwt.EmailFromToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorBody{ErrorMsg: aws.String(err.Error())})
		return
	}
	token := jwt.GetTokenRequest(c)
	// fmt.Println(token)
	if !usecase.TokenIsValid(email, token) {
		c.JSON(http.StatusBadRequest, domain.ErrorBody{ErrorMsg: aws.String(ErrorTokenInvalid)})
		return
	}
	getCharactersUsecase := usecase.GetCharactersUseCase{}
	result, err := getCharactersUsecase.Execute(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorBody{ErrorMsg: aws.String(err.Error())})
		return
	}

	c.JSON(200, result)
}
