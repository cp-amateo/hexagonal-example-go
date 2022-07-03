package main

import (
	"github.com/cp-amateo/hexagonal-example-go/adapters/mysql-repository"
	"github.com/cp-amateo/hexagonal-example-go/domain/service"
	"github.com/cp-amateo/hexagonal-example-go/rest-api/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func main() {
	db := configureDb()
	startWebEngine(db)
}

func startWebEngine(db *gorm.DB) {
	webEngine := gin.Default()

	customerRepository := mysql_repository.NewCustomerRepository(db)
	customerService := service.NewCustomerService(customerRepository)
	customerHandler := handler.NewCustomerHandler(customerService)

	webEngine.GET("/customer/:id", customerHandler.GetCustomerById)
	webEngine.POST("/customer", customerHandler.CreateCustomer)

	err := webEngine.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}

func configureDb() *gorm.DB {
	url := os.Getenv("MYSQL_URL")
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASS")
	schema := os.Getenv("MYSQL_SCHEMA")

	dsn := user + ":" + pass + "@tcp(" + url + ")/" + schema + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
