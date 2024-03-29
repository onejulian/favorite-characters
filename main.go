package main

import (
	"favorite-characters/src/infraestructure/jwt"
	"favorite-characters/src/view"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda
var r *gin.Engine

func init() {
	r = gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://characters.juliandavid.co", "http://localhost:3098", "https://characters-stage.juliandavid.co"}
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
	if runningInLambda() {
		lambda.Start(handler)
	} else {
		if r != nil {
			r.Run(":8080")
		}
	}
}

func runningInLambda() bool {
	_, exists := os.LookupEnv("AWS_LAMBDA_RUNTIME_API")
	return exists
}
