package controller

import (
	"github.com/AsrofunNiam/technical-tes-digdaya-olah-teknologi-indonesia/auth"
	"github.com/gin-gonic/gin"
)

type ProductController interface {
	FindAll(context *gin.Context, auth *auth.AccessDetails)
	FindByID(context *gin.Context, auth *auth.AccessDetails)
	FindImage(context *gin.Context, auth *auth.AccessDetails)

	// Group transaction
	FindAllProductTransaction(context *gin.Context, auth *auth.AccessDetails)
	CreateProductTransaction(context *gin.Context, auth *auth.AccessDetails)
}
