package s3

import (
	"bytes"
	"context"
	"fmt"
	"kost/configs"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/labstack/gommon/log"
)

type S3Client struct {
	s3 *s3.Client
}

func NewS3Client(config *configs.AppConfig) *S3Client {
	awsConfig, err := awsConfig.LoadDefaultConfig(context.TODO(),
		awsConfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			configs.Get().AwsS3.AccessKey,
			configs.Get().AwsS3.SecretKey,
			"",
		)),
		awsConfig.WithRegion(configs.Get().AwsS3.Region),
	)
	if err != nil {
		log.Warn(err)
	}
	client := s3.NewFromConfig(awsConfig)
	return &S3Client{
		s3: client,
	}
}

func (s *S3Client) UploadFileToS3(filename string, file multipart.FileHeader) (string, error) {
	// s3 Client
	uploader := manager.NewUploader(s.s3)
	src, err := file.Open()
	if err != nil {
		log.Warn(err)
		return "", err
	}
	defer src.Close()
	buffer := make([]byte, file.Size)
	src.Read(buffer)
	body, _ := file.Open()
	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(configs.Get().AwsS3.Bucket),
		ContentType: aws.String(http.DetectContentType(buffer)),
		Key:         aws.String(filename),
		Body:        body,
	})
	if err != nil {
		log.Warn(err)
		return "", err
	}
	return result.Location, nil
}

func (s *S3Client) DeleteFromS3(filename string) error {

	// s3 Client
	_, err := s.s3.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(configs.Get().AwsS3.Bucket),
		Key:    aws.String(filename),
	})
	if err != nil {
		log.Warn(err)
		return err
	}
	return nil
}

func (s *S3Client) UploadInvoiceToS3(filename string, url string) (string, error) {
	// s3 Client
	uploader := manager.NewUploader(s.s3)
	dir, err := os.Getwd()

	file, err := os.Open(fmt.Sprintf("%s/%s", dir, url))
	file.Close()

	if err != nil {
		log.Warn(err)
		return "", err
	}

	upFileInfo, _ := file.Stat()
	var fileSize int64 = upFileInfo.Size()
	fileBuffer := make([]byte, fileSize)
	file.Read(fileBuffer)

	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(configs.Get().AwsS3.Bucket),
		ContentType: aws.String(http.DetectContentType(fileBuffer)),
		Key:         aws.String(filename),
		Body:        bytes.NewReader(fileBuffer),
	})
	if err != nil {
		log.Warn(err)
		return "", err
	}
	err = os.Remove(fmt.Sprintf("%s/%s", dir, url))
	if err != nil {
		log.Warn(err)
		return "", err
	}

	return result.Location, nil
}
