package service

import (
	"github.com/cp-amateo/hexagonal-example-go/domain/model"
	"github.com/cp-amateo/hexagonal-example-go/domain/port"
)

type CustomerService struct {
	customerRepository port.CustomerRepository
}

func NewCustomerService(customerRepository port.CustomerRepository) *CustomerService {
	return &CustomerService{customerRepository}
}

func (cs *CustomerService) GetCustomerById(id int) (*model.Customer, error) {
	return cs.customerRepository.GetCustomerById(id)
}

func (cs *CustomerService) Save(customer model.Customer) (model.Customer, error) {
	return cs.customerRepository.Save(customer)
}
