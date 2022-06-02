package s3

import "mime/multipart"

type S3Control interface {
	UploadFileToS3(filename string, file multipart.FileHeader) (string, error)
	DeleteFromS3(filename string) error
}
