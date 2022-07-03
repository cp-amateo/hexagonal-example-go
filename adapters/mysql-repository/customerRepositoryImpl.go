package mysql_repository

import (
	"github.com/cp-amateo/hexagonal-example-go/adapters/mysql-repository/entity"
	"github.com/cp-amateo/hexagonal-example-go/domain/model"
	"github.com/dranikpg/dto-mapper"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{db: db}
}

func (cr *CustomerRepository) GetCustomerById(id int) (*model.Customer, error) {
	var customerEntity entity.Customer

	if err := cr.db.First(&customerEntity, "id = ?", id).Error; err != nil {
		return nil, err
	}

	customer, err := mapToDomain(customerEntity)
	return &customer, err
}

func (cr *CustomerRepository) Save(customer model.Customer) (model.Customer, error) {
	customerEntity, err := mapToEntity(customer)
	if err != nil {
		return customer, err
	}

	if err := cr.db.Create(&customerEntity).Error; err != nil {
		return customer, err
	}

	return mapToDomain(customerEntity)
}

func mapToDomain(customerEntity entity.Customer) (model.Customer, error) {
	var customer model.Customer
	err := dto.Map(&customer, customerEntity)

	return customer, err
}

func mapToEntity(customer model.Customer) (entity.Customer, error) {
	var customerEntity entity.Customer
	err := dto.Map(&customerEntity, customer)

	return customerEntity, err
}
