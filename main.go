package main

import (
	"mbs-back/src/infraestructure/jwt"
	"mbs-back/src/view"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

func init() {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://characters.juliandavid.me", "http://localhost:3098", "https://characters-stage.juliandavid.me"}
	config.AllowHeaders = append(config.AllowHeaders, "Authorization")

	r.Use(cors.New(config))

	r.POST("/register", view.CreateUser)
	r.POST("/login", view.Login)

	routes := r.Group("/", jwt.AuthorizeJWT())
	routes.GET("/logout", view.Logout)
	routes.GET("/users", view.GetUser)
	routes.PUT("/users", view.UpdateUser)
	routes.PUT("/users/password", view.ChangePassword)
	routes.DELETE("/users", view.DeleteUser)
	routes.POST("/characters", view.AddCharacter)
	routes.DELETE("/characters", view.DeleteCharacter)
	routes.GET("/characters", view.GetCharacters)

	ginLambda = ginadapter.New(r)
}

func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.Proxy(req)
}

func main() {
	lambda.Start(handler)
}
