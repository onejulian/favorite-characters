package jwt

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type ResponseLogin struct {
	Email string
	Token string
	Code  int
}

var SECRET = os.Getenv("SECRET")

func CreateJWT(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Minute * 120).Unix()

	tokenString, err := token.SignedString([]byte(SECRET))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func CheckPasswordHash(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	return err == nil
}

func getEmailFromToken(tokenString string) string {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET), nil
	})

	if err != nil {
		return ""
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		email := claims["email"].(string)
		return email
	}

	return ""
}

func GetTokenRequest(c *gin.Context) string {
	const BEARER_SCHEMA = "Bearer "
	authHeader := c.GetHeader("Authorization")
	if len(authHeader) == 0 {
		return ""
	}
	tokenString := strings.TrimSpace(authHeader[len(BEARER_SCHEMA):])
	return tokenString
}

func ValidateJWT(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET), nil
	})

	if err != nil {
		return false, err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true, nil
	}

	return false, nil
}

func EmailFromToken(c *gin.Context) (string, error) {
	email := getEmailFromToken(GetTokenRequest(c))
	if len(email) == 0 {
		return "", fmt.Errorf("invalid token when getting email")
	}

	return email, nil
}

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := GetTokenRequest(c)
		if len(token) == 0 {
			c.AbortWithStatusJSON(401, gin.H{"error": "token not found"})
			return
		}

		isValid, err := ValidateJWT(token)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			return
		}

		if !isValid {
			c.AbortWithStatusJSON(401, gin.H{"error": "invalid token"})
			return
		}

		c.Next()
	}
}

func HashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(hash)
}
