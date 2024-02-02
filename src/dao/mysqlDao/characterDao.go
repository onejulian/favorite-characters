package mysqldao

import (
	"errors"
	"favorite-characters/src/domain"
	"favorite-characters/src/infraestructure/constants"
	dbconfig "favorite-characters/src/infraestructure/dbconfig/mysqlconfig"

	"gorm.io/gorm"
)

type CharacterDao struct {
	db *gorm.DB
}

func NewCharacterDao() *CharacterDao {
	db := dbconfig.Db
	db.AutoMigrate(&domain.Character{})
	return &CharacterDao{db: db}
}

func (c *CharacterDao) Create(character domain.Character) (*domain.Character, error) {
	var count int64
	c.db.Model(&domain.Character{}).Where("user_email = ? AND id_character = ?", character.UserEmail, character.IdCharacter).Count(&count)
	if count > 0 {
		return nil, errors.New(constants.ErrorCharacterAlreadyExists)
	}

	if result := c.db.Create(&character); result.Error != nil {
		return nil, result.Error
	}

	return &character, nil
}

func (c *CharacterDao) Delete(userEmail string, idCharacter string) error {
	if result := c.db.Where("user_email = ? AND id_character = ?", userEmail, idCharacter).Delete(&domain.Character{}); result.Error != nil {
		return result.Error
	}

	return nil
}

func (c *CharacterDao) FindByEmail(userEmail string) ([]*domain.Character, error) {
	var characters []*domain.Character
	if result := c.db.Where("user_email = ?", userEmail).Find(&characters); result.Error != nil {
		return nil, result.Error
	}

	return characters, nil
}

func (c *CharacterDao) CharacterExists(userEmail string, idCharacter string) (bool, error) {
	var count int64
	c.db.Model(&domain.Character{}).Where("user_email = ? AND id_character = ?", userEmail, idCharacter).Count(&count)
	return count > 0, nil
}

func (c *CharacterDao) DeleteAll(userEmail string) error {
	if result := c.db.Where("user_email = ?", userEmail).Delete(&domain.Character{}); result.Error != nil {
		return result.Error
	}

	return nil
}
