package handler

import (
	"goshop/internal/repository"

	"gorm.io/gorm"
)

type UserRepository struct {
	Repo repository.UserRepo
}

func UserControllerInit(db *gorm.DB) *UserRepository {
	return &UserRepository{
		Repo: repository.UserRepoIntf(db),
	}
}
