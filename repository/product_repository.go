package repository

import (
	"github.com/AsrofunNiam/technical-tes-digdaya-olah-teknologi-indonesia/model/domain"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll(db *gorm.DB, filters *map[string]string) domain.Products
	FindByID(db *gorm.DB, id *uint) domain.Product
}
