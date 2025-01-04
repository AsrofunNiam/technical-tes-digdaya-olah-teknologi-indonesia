package repository

import (
	"github.com/AsrofunNiam/technical-tes-digdaya-olah-teknologi-indonesia/helper"
	domain "github.com/AsrofunNiam/technical-tes-digdaya-olah-teknologi-indonesia/model/domain"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Login(db *gorm.DB, identity *string) domain.User {
	var user domain.User
	err := db.Where("email = ? OR number_phone = ?", identity, identity).First(&user).Error
	helper.PanicIfError(err)
	return user
}
