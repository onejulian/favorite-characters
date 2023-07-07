package view

import (
	"mbs-back/src/domain"
	"mbs-back/src/infraestructure/jwt"
	"mbs-back/src/usecase"
	"mbs-back/src/view/req"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/gin-gonic/gin"
)

var (
	ErrorTokenInvalid = "el token es invalido"
	SuccesLogout      = "la sesion se cerro correctamente"
	SuccesPassword    = "la contraseña se cambio correctamente"
)

func Login(c *gin.Context) {
	loginRequest := req.LoginRequest{}
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorBody{ErrorMsg: aws.String(err.Error())})
		return
	}

	loginUseCase := usecase.LoginUseCase{}

	resp, err := loginUseCase.Execute(loginRequest.Email, loginRequest.Password)
	if err != nil {
		c.JSON(resp.Code, domain.ErrorBody{ErrorMsg: aws.String(err.Error())})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func Logout(c *gin.Context) {
	logoutUseCase := usecase.LogoutUseCase{}

	token := jwt.GetTokenRequest(c)
	email, err := jwt.EmailFromToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorBody{ErrorMsg: aws.String(err.Error())})
		return
	}

	if !usecase.TokenIsValid(email, token) {
		c.JSON(http.StatusUnauthorized, domain.ErrorBody{ErrorMsg: aws.String(ErrorTokenInvalid)})
		return
	}

	tokenStruct := domain.Token{Value: token, UserEmail: email}

	code, err := logoutUseCase.Execute(tokenStruct)
	if err != nil {
		c.JSON(code, domain.ErrorBody{ErrorMsg: aws.String(err.Error())})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessBody{SuccessMsg: aws.String(SuccesLogout)})
}

func ChangePassword(c *gin.Context) {
	req := req.ChangePasswordRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorBody{ErrorMsg: aws.String(err.Error())})
		return
	}
	token := jwt.GetTokenRequest(c)
	if !usecase.TokenIsValid(req.Email, token) {
		c.JSON(http.StatusBadRequest, domain.ErrorBody{ErrorMsg: aws.String(ErrorTokenInvalid)})
		return
	}

	changePasswordUseCase := usecase.ChangePasswordUseCase{}
	code, err := changePasswordUseCase.Execute(req.Email, req.NewPassword, req.OldPassword)
	if err != nil {
		c.JSON(code, domain.ErrorBody{ErrorMsg: aws.String(err.Error())})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessBody{SuccessMsg: aws.String(SuccesPassword)})
}
