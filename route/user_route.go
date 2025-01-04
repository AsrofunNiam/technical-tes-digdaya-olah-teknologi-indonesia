package route

import (
	"github.com/AsrofunNiam/technical-tes-digdaya-olah-teknologi-indonesia/controller"
	"github.com/AsrofunNiam/technical-tes-digdaya-olah-teknologi-indonesia/repository"
	"github.com/AsrofunNiam/technical-tes-digdaya-olah-teknologi-indonesia/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func UserRoute(router *gin.Engine, db *gorm.DB, validate *validator.Validate) {

	userService := service.NewUserService(
		repository.NewUserRepository(),
		db,
		validate)
	userController := controller.NewUserController(userService)
	router.POST("/users/login", userController.Login)
}
