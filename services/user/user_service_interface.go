package user

import (
	"kost/entities"
	storageProvider "kost/services/storage"
	"mime/multipart"
)

type UserServiceInterface interface {
	CreateUser(internalRequest entities.CreateUserRequest, files map[string]*multipart.FileHeader, storageProvider storageProvider.StorageInterface) (entities.InternalAuthResponse, error)
	GetCustomer(id int) (entities.CustomerResponse, error)
	GetAllMember() ([]entities.User, error)
	GetInternal(id int) (entities.InternalResponse, error)
	UpdateInternal(customerRequest entities.UpdateInternalRequest, id int, files map[string]*multipart.FileHeader, storageProvider storageProvider.StorageInterface) (entities.InternalResponse, error)
	UpdateCustomer(customerRequest entities.UpdateCustomerRequest, id int, files map[string]*multipart.FileHeader, storage storageProvider.StorageInterface) (entities.CustomerResponse, error)
	DeleteInternal(id int, storage storageProvider.StorageInterface) error
	DeleteCustomer(id int, storage storageProvider.StorageInterface) error
}
