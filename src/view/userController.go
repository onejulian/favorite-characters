package view

import (
	"mbs-back/src/domain"
	"mbs-back/src/infraestructure/jwt"
	"mbs-back/src/usecase"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/gin-gonic/gin"
)

var (
	ErrorMethodNotAllowed string = "Método no permitido"
	ErrorVoidEmail        string = "Email no puede estar vacío"
	ErrorPermissionDenied string = "No tiene permisos para acceder a este recurso"
)

func GetUser(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorBody{ErrorMsg: aws.String(ErrorVoidEmail)})
		return
	}

	emailFromToken, err := jwt.EmailFromToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorBody{ErrorMsg: aws.String(err.Error())})
		return
	}
	token := jwt.GetTokenRequest(c)
	if !usecase.TokenIsValid(email, token) {
		c.JSON(http.StatusBadRequest, domain.ErrorBody{ErrorMsg: aws.String(ErrorTokenInvalid)})
		return
	}

	if emailFromToken != email {
		c.JSON(http.StatusUnauthorized, domain.ErrorBody{ErrorMsg: aws.String(ErrorPermissionDenied)})
		return
	}

	getUserUsecase := usecase.GetUserUseCase{}
	result, err := getUserUsecase.Execute(email)
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrorBody{ErrorMsg: aws.String(err.Error())})
		return
	}
	c.JSON(http.StatusOK, result)
}

func CreateUser(c *gin.Context) {
	user := domain.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorBody{ErrorMsg: aws.String(err.Error())})
		return
	}
	createUserUsecase := usecase.CreateUserUseCase{}
	result, err := createUserUsecase.Execute(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorBody{ErrorMsg: aws.String(err.Error())})
		return
	}
	c.JSON(http.StatusCreated, result)
}

func UpdateUser(c *gin.Context) {
	user := domain.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorBody{ErrorMsg: aws.String(err.Error())})
		return
	}

	emailFromToken, err := jwt.EmailFromToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorBody{ErrorMsg: aws.String(err.Error())})
		return
	}
	token := jwt.GetTokenRequest(c)
	if !usecase.TokenIsValid(emailFromToken, token) {
		c.JSON(http.StatusBadRequest, domain.ErrorBody{ErrorMsg: aws.String(ErrorTokenInvalid)})
		return
	}

	if emailFromToken != user.Email {
		c.JSON(http.StatusUnauthorized, domain.ErrorBody{ErrorMsg: aws.String(ErrorPermissionDenied)})
		return
	}

	updateUserUsecase := usecase.UpdateUserUseCase{}
	result, err := updateUserUsecase.Execute(user, token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorBody{ErrorMsg: aws.String(err.Error())})
		return
	}
	c.JSON(http.StatusOK, result)
}

func DeleteUser(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorBody{ErrorMsg: aws.String(ErrorVoidEmail)})
		return
	}

	emailFromToken, err := jwt.EmailFromToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorBody{ErrorMsg: aws.String(err.Error())})
		return
	}
	token := jwt.GetTokenRequest(c)
	if !usecase.TokenIsValid(emailFromToken, token) {
		c.JSON(http.StatusBadRequest, domain.ErrorBody{ErrorMsg: aws.String(ErrorTokenInvalid)})
		return
	}

	if emailFromToken != email {
		c.JSON(http.StatusUnauthorized, domain.ErrorBody{ErrorMsg: aws.String(ErrorPermissionDenied)})
		return
	}

	deleteUserUsecase := usecase.DeleteUserUseCase{}
	err = deleteUserUsecase.Execute(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorBody{ErrorMsg: aws.String(err.Error())})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado correctamente"})
}
