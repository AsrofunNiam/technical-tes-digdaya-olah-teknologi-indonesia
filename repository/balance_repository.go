package repository

import (
	"github.com/AsrofunNiam/technical-tes-digdaya-olah-teknologi-indonesia/model/domain"
	"gorm.io/gorm"
)

type BalanceRepository interface {
	FindAll(db *gorm.DB, filters *map[string]string) domain.Balances
	FindByID(db *gorm.DB, userID *uint) domain.Balance
	Update(db *gorm.DB, limit *domain.Balance) *domain.Balance
}
