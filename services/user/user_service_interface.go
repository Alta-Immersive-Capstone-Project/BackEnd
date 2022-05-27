package user

import (
	"backend/be8/entities"
	storageProvider "backend/be8/services/storage"
	"mime/multipart"
)

type UserServiceInterface interface {
	CreateUser(internalRequest entities.CreateUserRequest, files map[string]*multipart.FileHeader, storageProvider storageProvider.StorageInterface) (entities.InternalAuthResponse, error)
	GetCustomer(id int) (entities.CustomerResponse, error)
	GetInternal(id int) (entities.InternalResponse, error)
	UpdateInternal(customerRequest entities.UpdateInternalRequest, id int, files map[string]*multipart.FileHeader, storageProvider storageProvider.StorageInterface) (entities.InternalResponse, error)
	DeleteInternal(id int, storage storageProvider.StorageInterface) error
	DeleteCustomer(id int, storage storageProvider.StorageInterface) error
}
