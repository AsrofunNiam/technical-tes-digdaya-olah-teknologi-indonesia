package app

import (
	"log"
	"os"
	"time"

	"github.com/AsrofunNiam/technical-tes-digdaya-olah-teknologi-indonesia/model/domain"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDatabase(user, host, password, port, db string) *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + db + "?parseTime=true"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}

	//  function auto migrate, create and generate schema table
	err = database.AutoMigrate(
		&domain.User{},
		&domain.Balance{},

		&domain.Transaction{},

		&domain.Product{},
		&domain.Currency{},
		&domain.ProductPrice{},
		&domain.Company{},
	)
	if err != nil {
		panic("failed to auto migrate schema")

	}

	return database
}

func ConnectClientCRedis(host, port, password string) *redis.Client {

	rdb := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       1,
		Protocol: 3,
	})

	return rdb

}
