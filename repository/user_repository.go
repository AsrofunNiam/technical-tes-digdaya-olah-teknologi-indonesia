package repository

import (
	domain "github.com/AsrofunNiam/technical-tes-digdaya-olah-teknologi-indonesia/model/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	Login(db *gorm.DB, identity *string) domain.User
}
