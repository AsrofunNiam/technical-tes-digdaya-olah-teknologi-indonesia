package service

import (
	"github.com/AsrofunNiam/technical-tes-digdaya-olah-teknologi-indonesia/auth"
	"github.com/AsrofunNiam/technical-tes-digdaya-olah-teknologi-indonesia/model/web"
	"github.com/gin-gonic/gin"
)

type ProductService interface {
	FindAll(auth *auth.AccessDetails, filters *map[string]string, c *gin.Context) []web.ProductResponse
	FindByID(auth *auth.AccessDetails, id *uint, c *gin.Context) web.ProductResponse
	FindImage(auth *auth.AccessDetails, image string, c *gin.Context) []byte

	// Group transaction
	FindAllTransaction(auth *auth.AccessDetails, filters *map[string]string, c *gin.Context) []web.TransactionResponse
	CreateTransaction(auth *auth.AccessDetails, request *web.TransactionCreateRequest, c *gin.Context) web.TransactionResponse
}
