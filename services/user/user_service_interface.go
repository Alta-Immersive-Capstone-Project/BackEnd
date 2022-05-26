package user

import (
	"backend/be8/entities"
	storageProvider "backend/be8/services/storage"
	"mime/multipart"
)

type UserServiceInterface interface {
	CreateInternal(internalRequest entities.CreateInternalRequest, files map[string]*multipart.FileHeader, storageProvider storageProvider.StorageInterface) (entities.InternalAuthResponse, error)
}
