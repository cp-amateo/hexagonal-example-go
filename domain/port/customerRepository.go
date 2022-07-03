package port

import "github.com/cp-amateo/hexagonal-example-go/domain/model"

type CustomerRepository interface {
	GetCustomerById(id int) (*model.Customer, error)
	Save(customer model.Customer) (model.Customer, error)
}
