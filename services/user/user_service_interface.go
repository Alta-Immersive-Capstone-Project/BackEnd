package user

import (
	"kost/entities"
)

type UserServiceInterface interface {
	CreateUser(internalRequest entities.CreateUserRequest, url string) (entities.User, error)
	GetbyID(id uint) (entities.CustomerResponse, error)
	GetAllMember() ([]entities.GetAll, error)
	UpdateInternal(customerRequest entities.UpdateInternalRequest, id uint, url string) (entities.InternalResponse, error)
	UpdateCustomer(customerRequest entities.UpdateCustomerRequest, id uint, url string) (entities.CustomerResponse, error)
	DeleteInternal(id uint) error
	DeleteCustomer(id uint) error
}
