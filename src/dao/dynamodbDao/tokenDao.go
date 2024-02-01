package dynamodbdao

import (
	"favorite-characters/src/domain"
	dbconfig "favorite-characters/src/infraestructure/dbconfig/dynamo"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type TokenDao struct {
	tableName  string
	dynaClient dynamodbiface.DynamoDBAPI
}

func NewTokenDao(tableName string) *TokenDao {
	return &TokenDao{
		tableName:  tableName,
		dynaClient: dbconfig.Client,
	}
}

func (t *TokenDao) Create(token domain.Token) (*domain.Token, error) {
	input := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"user_email": {
				S: aws.String(token.UserEmail),
			},
			"token": {
				S: aws.String(token.Value),
			},
		},
		TableName: aws.String(t.tableName),
	}

	_, err := t.dynaClient.PutItem(input)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (t *TokenDao) Delete(token domain.Token) error {
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"token": {
				S: aws.String(token.Value),
			},
			"user_email": {
				S: aws.String(token.UserEmail),
			},
		},
		TableName: aws.String(t.tableName),
	}

	_, err := t.dynaClient.DeleteItem(input)
	if err != nil {
		return err
	}

	return nil
}

func (t *TokenDao) FindByEmail(userEmail string) (*[]domain.Token, error) {
	resp, err := t.dynaClient.Query(&dynamodb.QueryInput{
		TableName: aws.String(t.tableName),
		KeyConditions: map[string]*dynamodb.Condition{
			"user_email": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(userEmail),
					},
				},
			},
		},
	})

	if err != nil {
		return nil, err
	}

	tokens := []domain.Token{}
	for _, item := range resp.Items {
		token := domain.Token{
			UserEmail: userEmail,
			Value:     *item["token"].S,
		}
		tokens = append(tokens, token)
	}

	return &tokens, nil
}

func (t *TokenDao) DeleteAll(userEmail string) error {
	tokens, err := t.FindByEmail(userEmail)
	if err != nil {
		return err
	}

	for _, token := range *tokens {
		err = t.Delete(token)
		if err != nil {
			return err
		}
	}

	return nil
}

func (t *TokenDao) ValidateToken(token, email string) (bool, error) {
	tokens, err := t.FindByEmail(email)
	if err != nil {
		return false, err
	}

	for _, tokenObj := range *tokens {
		if tokenObj.Value == token {
			return true, nil
		}
	}

	return false, nil
}
