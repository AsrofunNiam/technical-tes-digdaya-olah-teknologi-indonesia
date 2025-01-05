package service

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/AsrofunNiam/technical-tes-digdaya-olah-teknologi-indonesia/auth"
	"github.com/AsrofunNiam/technical-tes-digdaya-olah-teknologi-indonesia/exception"
	"github.com/AsrofunNiam/technical-tes-digdaya-olah-teknologi-indonesia/helper"
	"github.com/AsrofunNiam/technical-tes-digdaya-olah-teknologi-indonesia/model/domain"
	"github.com/AsrofunNiam/technical-tes-digdaya-olah-teknologi-indonesia/model/web"
	"github.com/AsrofunNiam/technical-tes-digdaya-olah-teknologi-indonesia/repository"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ProductServiceImpl struct {
	ProductRepository     repository.ProductRepository
	TransactionRepository repository.TransactionRepository
	BalanceRepository     repository.BalanceRepository
	DB                    *gorm.DB
	Validate              *validator.Validate
}

func NewProductService(
	productRepository repository.ProductRepository,
	transactionRepository repository.TransactionRepository,
	limitRepository repository.BalanceRepository,
	db *gorm.DB,
	validate *validator.Validate,
) ProductService {
	return &ProductServiceImpl{
		ProductRepository:     productRepository,
		TransactionRepository: transactionRepository,
		BalanceRepository:     limitRepository,
		DB:                    db,
		Validate:              validate,
	}
}

func (service *ProductServiceImpl) FindAll(auth *auth.AccessDetails, filters *map[string]string, c *gin.Context) []web.ProductResponse {
	tx := service.DB
	err := tx.Error
	helper.PanicIfError(err)

	products := service.ProductRepository.FindAll(tx, filters)
	return products.ToProductResponses()
}

func (service *ProductServiceImpl) FindByID(auth *auth.AccessDetails, id *uint, c *gin.Context) web.ProductResponse {
	tx := service.DB
	err := tx.Error
	helper.PanicIfError(err)

	product := service.ProductRepository.FindByID(tx, id)
	return product.ToProductResponse()
}

func (service *ProductServiceImpl) FindImage(auth *auth.AccessDetails, imagesName string, c *gin.Context) []byte {
	// Open file
	fileProofPhoto := helper.PathToProduct + imagesName
	fileOpen, err := os.Open(fileProofPhoto)
	helper.PanicIfError(err)

	fileBytes, err := io.ReadAll(fileOpen)
	helper.PanicIfError(err)

	return fileBytes
}

// Group tRansaction Product

func (service *ProductServiceImpl) FindAllTransaction(auth *auth.AccessDetails, filters *map[string]string, c *gin.Context) []web.TransactionResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	transaction := service.TransactionRepository.FindAll(tx, filters)
	return transaction.ToTransactionResponses()
}

func (service *ProductServiceImpl) CreateTransaction(auth *auth.AccessDetails, request *web.TransactionCreateRequest, c *gin.Context) web.TransactionResponse {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	//  Validate request
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	channel := make(chan domain.Product)
	defer close(channel)

	// Validate user role
	if auth.Role != "customer" {
		err := &exception.ErrorSendToResponse{Err: "Only customers can create loans"}
		helper.PanicIfError(err)
	}

	// Find product
	// product := service.ProductRepository.FindByID(tx, &request.ProductID)

	// implement chanel to find product
	go func() {
		product := service.ProductRepository.FindByID(tx, &request.ProductID)
		channel <- product
		fmt.Println("done query product")

	}()

	product := <-channel

	// Find limit
	limit := service.BalanceRepository.FindByID(tx, &auth.ID)

	// Validate limit existence
	if limit.ID == 0 {
		err := &exception.ErrorSendToResponse{Err: "Balance not found"}
		helper.PanicIfError(err)
	}

	totalValue := product.ProductPrice.Price + request.AdminFee

	// Validate limit against instalment amount
	if limit.Value < totalValue {
		err := &exception.ErrorSendToResponse{Err: "Insufficient limit for the requested product"}
		helper.PanicIfError(err)
	}

	transaction := &domain.Transaction{
		// Required Fields
		CreatedByID: auth.ID,
		UserID:      auth.ID,

		// Fields for Loan
		OnTheRoad:       product.ProductPrice.Price,
		AdminFee:        request.AdminFee,
		TotalValue:      totalValue,
		ProductID:       request.ProductID,
		ProductName:     product.Name,
		TransactionType: request.TransactionType,
		NumberContract:  fmt.Sprintf("%d-%d-%d", auth.ID, product.ID, time.Now().Unix()),
	}

	transaction = service.TransactionRepository.Create(tx, transaction)

	// Update limit
	limit.Value -= totalValue
	service.BalanceRepository.Update(tx, &limit)

	return transaction.ToTransactionResponse()
}
