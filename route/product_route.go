package route

import (
	"github.com/AsrofunNiam/technical-tes-digdaya-olah-teknologi-indonesia/auth"
	"github.com/AsrofunNiam/technical-tes-digdaya-olah-teknologi-indonesia/controller"
	"github.com/AsrofunNiam/technical-tes-digdaya-olah-teknologi-indonesia/repository"
	"github.com/AsrofunNiam/technical-tes-digdaya-olah-teknologi-indonesia/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func ProductRoute(router *gin.Engine, db *gorm.DB, validate *validator.Validate) {
	Products := service.NewProductService(
		repository.NewProductRepository(),
		db,
		validate,
	)
	productController := controller.NewProductController(Products)
	router.GET("/products", auth.Auth(productController.FindAll, []string{}))
	router.GET("/products/:id", auth.Auth(productController.FindByID, []string{}))
}