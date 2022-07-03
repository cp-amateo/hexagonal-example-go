package handler

import (
	"github.com/cp-amateo/hexagonal-example-go/domain/model"
	"github.com/cp-amateo/hexagonal-example-go/domain/service"
	"github.com/cp-amateo/hexagonal-example-go/rest-api/dto"
	mapper "github.com/dranikpg/dto-mapper"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CustomerHandler struct {
	CustomerService *service.CustomerService
}

func NewCustomerHandler(customerService *service.CustomerService) *CustomerHandler {
	return &CustomerHandler{CustomerService: customerService}
}

func (ch *CustomerHandler) GetCustomerById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var customer *model.Customer
	var err error

	if customer, err = ch.CustomerService.GetCustomerById(id); err != nil {
		c.IndentedJSON(http.StatusNotFound, nil)
		return
	}

	var customerDto = dto.Customer{}
	if err = mapper.Map(&customerDto, customer); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	c.IndentedJSON(http.StatusOK, customerDto)
}

func (ch *CustomerHandler) CreateCustomer(c *gin.Context) {
	var newCustomerDto dto.Customer
	var err error

	if err = c.BindJSON(&newCustomerDto); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	var newCustomer model.Customer
	if err = mapper.Map(&newCustomer, newCustomerDto); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	var customerSaved model.Customer
	if customerSaved, err = ch.CustomerService.Save(newCustomer); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	var customerSavedDto = dto.Customer{}
	if err = mapper.Map(&customerSavedDto, customerSaved); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	c.IndentedJSON(http.StatusOK, customerSavedDto)
}
