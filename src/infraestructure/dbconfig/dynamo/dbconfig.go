package dbconfig

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var Client = initDynaClient()

func initDynaClient() dynamodbiface.DynamoDBAPI {
	region := os.Getenv("AWS_REGION")
	is_local := os.Getenv("LOCAL")

	config := &aws.Config{
		Region: aws.String(region),
	}

	if is_local == "1" {
		awsAccessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
		awsSecretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
		config.Credentials = credentials.NewStaticCredentials(awsAccessKeyID, awsSecretAccessKey, "")
	}

	awsSession, err := session.NewSession(config)
	if err != nil {
		return nil
	}

	return dynamodb.New(awsSession)
}
