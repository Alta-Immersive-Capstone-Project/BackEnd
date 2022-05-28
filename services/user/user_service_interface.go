package user

import (
	"kost/entities"
	storageProvider "kost/services/storage"
	"mime/multipart"
)

type UserServiceInterface interface {
	CreateInternal(internalRequest entities.CreateInternalRequest, files map[string]*multipart.FileHeader, storageProvider storageProvider.StorageInterface) (entities.InternalAuthResponse, error)
}
