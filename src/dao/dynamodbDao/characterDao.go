package dynamodbdao

import (
	"errors"
	"favorite-characters/src/domain"
	"favorite-characters/src/infraestructure/constants"
	dbconfig "favorite-characters/src/infraestructure/dbconfig/dynamoconfig"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type CharacterDao struct {
	tableName  string
	dynaClient dynamodbiface.DynamoDBAPI
}

func NewCharacterDao(tableName string) *CharacterDao {
	return &CharacterDao{
		tableName:  tableName,
		dynaClient: dbconfig.Client,
	}
}

func (c *CharacterDao) Create(character domain.Character) (*domain.Character, error) {
	currentCharacter, _ := c.CharacterExists(character.UserEmail, character.IdCharacter)
	if currentCharacter {
		return nil, errors.New(constants.ErrorCharacterAlreadyExists)
	}

	input := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"user_email": {
				S: aws.String(character.UserEmail),
			},
			"character_id": {
				S: aws.String(character.IdCharacter),
			},
		},
		TableName: aws.String(c.tableName),
	}

	_, err := c.dynaClient.PutItem(input)
	if err != nil {
		return nil, err
	}

	return &character, nil
}

func (c *CharacterDao) Delete(userEmail string, idCharacter string) error {
	currentCharacter, _ := c.CharacterExists(userEmail, idCharacter)
	if !currentCharacter {
		return errors.New(constants.ErrorCharacterNotFound)
	}

	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"user_email": {
				S: aws.String(userEmail),
			},
			"character_id": {
				S: aws.String(idCharacter),
			},
		},
		TableName: aws.String(c.tableName),
	}

	_, err := c.dynaClient.DeleteItem(input)
	if err != nil {
		return err
	}

	return nil
}

func (c *CharacterDao) FindByEmail(userEmail string) ([]*domain.Character, error) {
	resp, err := c.dynaClient.Query(&dynamodb.QueryInput{
		TableName: aws.String(c.tableName),
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

	characters := make([]*domain.Character, 0)
	for _, item := range resp.Items {
		character := domain.Character{
			UserEmail:   userEmail,
			IdCharacter: *item["character_id"].S,
		}
		characters = append(characters, &character)
	}

	return characters, nil
}

func (c *CharacterDao) CharacterExists(userEmail string, idCharacter string) (bool, error) {
	resp, err := c.dynaClient.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(c.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"user_email": {
				S: aws.String(userEmail),
			},
			"character_id": {
				S: aws.String(idCharacter),
			},
		},
	})
	if err != nil {
		return false, err
	}

	return resp.Item != nil, nil
}

func (c *CharacterDao) DeleteAll(userEmail string) error {
	characters, err := c.FindByEmail(userEmail)
	if err != nil {
		return err
	}

	for _, character := range characters {
		err := c.Delete(userEmail, character.IdCharacter)
		if err != nil {
			return err
		}
	}

	return nil
}
