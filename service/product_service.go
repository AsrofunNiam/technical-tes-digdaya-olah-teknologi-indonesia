package service

import (
	"github.com/AsrofunNiam/technical-tes-digdaya-olah-teknologi-indonesia/auth"
	"github.com/AsrofunNiam/technical-tes-digdaya-olah-teknologi-indonesia/model/web"
	"github.com/gin-gonic/gin"
)

type ProductService interface {
	FindAll(auth *auth.AccessDetails, filters *map[string]string, c *gin.Context) []web.ProductResponse
	FindByID(auth *auth.AccessDetails, id *uint, c *gin.Context) web.ProductResponse
}
