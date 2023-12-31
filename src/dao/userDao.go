package dao

import (
	"errors"
	"mbs-back/src/domain"
	"mbs-back/src/infraestructure/jwt"
	"mbs-back/src/infraestructure/util"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var (
	ErrorInvalidEmail      = "email inválido"
	ErrorUserAlreadyExists = "el usuario ya existe"
	ErrorUserNotFound      = "usuario no encontrado"
	client                 = initDynaClient()
)

type UserDao struct {
	tableName  string
	dynaClient dynamodbiface.DynamoDBAPI
}

func NewUserDao(tableName string) *UserDao {
	return &UserDao{
		tableName:  tableName,
		dynaClient: client,
	}
}

func initDynaClient() dynamodbiface.DynamoDBAPI {
	region := os.Getenv("AWS_REGION")

	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(region)})
	if err != nil {
		return nil
	}

	return dynamodb.New(awsSession)
}

func (u *UserDao) Create(user domain.User) (*domain.User, error) {
	user.IsActive = true
	if !util.IsEmailValid(user.Email) {
		return nil, errors.New(ErrorInvalidEmail)
	}

	currentUser, _ := u.FindByEmail(user.Email)
	if currentUser != nil {
		return nil, errors.New(ErrorUserAlreadyExists)
	}

	input := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"email": {
				S: aws.String(user.Email),
			},
			"firstName": {
				S: aws.String(user.FirtsName),
			},
			"lastName": {
				S: aws.String(user.LastName),
			},
			"password": {
				S: aws.String(jwt.HashAndSalt([]byte(user.Password))),
			},
			"is_active": {
				BOOL: aws.Bool(user.IsActive),
			},
		},
		TableName: aws.String(u.tableName),
	}

	_, err := u.dynaClient.PutItem(input)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserDao) Update(user domain.User) (*domain.User, error) {
	currentUser, _ := u.FindByEmail(user.Email)
	if currentUser == nil {
		return nil, errors.New(ErrorUserNotFound)
	}

	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":f": {
				S: aws.String(user.FirtsName),
			},
			":l": {
				S: aws.String(user.LastName),
			},
			":s": {
				BOOL: aws.Bool(user.IsActive),
			},
		},
		Key: map[string]*dynamodb.AttributeValue{
			"email": {
				S: aws.String(user.Email),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		TableName:        aws.String(u.tableName),
		UpdateExpression: aws.String("set firstName = :f, lastName = :l, is_active = :s"),
	}

	_, err := u.dynaClient.UpdateItem(input)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserDao) Delete(email string) error {
	currentUser, _ := u.FindByEmail(email)
	if currentUser == nil {
		return errors.New(ErrorUserNotFound)
	}

	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"email": {
				S: aws.String(email),
			},
		},
		TableName: aws.String(u.tableName),
	}

	_, err := u.dynaClient.DeleteItem(input)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserDao) FindByEmail(email string) (*domain.User, error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"email": {
				S: aws.String(email),
			},
		},
		TableName: aws.String(u.tableName),
	}

	result, err := u.dynaClient.GetItem(input)
	if err != nil {
		return nil, err
	}

	user := new(domain.User)
	err = dynamodbattribute.UnmarshalMap(result.Item, user)
	if err != nil {
		return nil, err
	}

	if len(user.Email) == 0 {
		return nil, errors.New(ErrorUserNotFound)
	}

	return user, nil
}

func (u *UserDao) ChangePassword(email string, password string) error {
	currentUser, _ := u.FindByEmail(email)
	if currentUser == nil {
		return errors.New(ErrorUserNotFound)
	}

	currentUser.Password = password
	
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":p": {
				S: aws.String(jwt.HashAndSalt([]byte(password))),
			},
		},
		Key: map[string]*dynamodb.AttributeValue{
			"email": {
				S: aws.String(email),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		TableName:        aws.String(u.tableName),
		UpdateExpression: aws.String("set password = :p"),
	}

	_, err := u.dynaClient.UpdateItem(input)
	if err != nil {
		return err
	}

	return nil
}
