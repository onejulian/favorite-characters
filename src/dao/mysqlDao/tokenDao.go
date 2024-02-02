package mysqldao

import (
	"errors"
	"favorite-characters/src/domain"
	dbconfig "favorite-characters/src/infraestructure/dbconfig/mysqlconfig"

	"gorm.io/gorm"
)

type TokenDao struct {
	db *gorm.DB
}

func NewTokenDao() *TokenDao {
	db := dbconfig.Db
	db.AutoMigrate(&domain.Token{})
	return &TokenDao{db: db}
}

func (t *TokenDao) Create(token domain.Token) (*domain.Token, error) {
	if result := t.db.Create(&token); result.Error != nil {
		return nil, result.Error
	}
	return &token, nil
}

func (t *TokenDao) Delete(token domain.Token) error {
	if result := t.db.Where("user_email = ? AND value = ?", token.UserEmail, token.Value).Delete(&domain.Token{}); result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *TokenDao) FindByEmail(userEmail string) (*[]domain.Token, error) {
	var tokens []domain.Token
	if result := t.db.Where("user_email = ?", userEmail).Find(&tokens); result.Error != nil {
		return nil, result.Error
	}
	return &tokens, nil
}

func (t *TokenDao) DeleteAll(userEmail string) error {
	if result := t.db.Where("user_email = ?", userEmail).Delete(&domain.Token{}); result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *TokenDao) ValidateToken(tokenValue, email string) (bool, error) {
	var token domain.Token
	if result := t.db.Where("user_email = ? AND value = ?", email, tokenValue).First(&token); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}
	return true, nil
}
