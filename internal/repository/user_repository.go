package repository

import (
	"goshop/database"
	models "goshop/internal/model"

	"gorm.io/gorm"
)

type UserDB struct {
	UDB *gorm.DB
}

type UserRepo interface {
	FindUser(models.User) (models.User, error)
	SignUp(models.User) error
}

func UserRepoIntf(db *gorm.DB) UserRepo {
	return &UserDB{
		UDB: database.DbClient,
	}
}

func (u *UserDB) FindUser(user models.User) (models.User, error) {
	err := u.UDB.Where("username = ?", user.Username).First(&user).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (u *UserDB) SignUp(user models.User) error {
	err := u.UDB.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}
