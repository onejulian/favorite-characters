package mysqldao

import (
	"errors"
	"favorite-characters/src/domain"
	"favorite-characters/src/infraestructure/constants"
	dbconfig "favorite-characters/src/infraestructure/dbconfig/mysqlconfig"
	"favorite-characters/src/infraestructure/jwt"
	"favorite-characters/src/infraestructure/util"

	"gorm.io/gorm"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao() *UserDao {
	db := dbconfig.Db
	db.AutoMigrate(&domain.User{})
	return &UserDao{db: db}
}

func (u *UserDao) Create(user domain.User) (*domain.User, error) {
	if !util.IsEmailValid(user.Email) {
		return nil, errors.New(constants.ErrorInvalidEmail)
	}

	var count int64
	u.db.Model(&domain.User{}).Where("email = ?", user.Email).Count(&count)
	if count > 0 {
		return nil, errors.New(constants.ErrorUserAlreadyExists)
	}

	user.Password = jwt.HashAndSalt([]byte(user.Password))
	user.IsActive = true

	if result := u.db.Create(&user); result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (u *UserDao) Update(user domain.User) (*domain.User, error) {
	result := u.db.Model(&domain.User{}).Where("email = ?", user.Email).Updates(domain.User{
		FirtsName: user.FirtsName,
		LastName:  user.LastName,
		IsActive:  user.IsActive,
	})
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New(constants.ErrorUserNotFound)
	}
	return &user, nil
}

func (u *UserDao) Delete(email string) error {
	result := u.db.Where("email = ?", email).Delete(&domain.User{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New(constants.ErrorUserNotFound)
	}
	return nil
}

func (u *UserDao) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	result := u.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New(constants.ErrorUserNotFound)
		}
		return nil, result.Error
	}
	return &user, nil
}

func (u *UserDao) ChangePassword(email string, newPassword string) error {
	hashedPassword := jwt.HashAndSalt([]byte(newPassword))
	result := u.db.Model(&domain.User{}).Where("email = ?", email).Update("password", hashedPassword)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New(constants.ErrorUserNotFound)
	}
	return nil
}
