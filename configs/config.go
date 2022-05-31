package configs

import (
	"fmt"
	"os"
	"sync"
)

type AppConfig struct {
	App struct {
		BaseURL string
		Port    string
		ENV     string
	}

	Database struct {
		Username string
		Password string
		Host     string
		Port     string
		Name     string
	}

	AwsS3 struct {
		Bucket    string
		Region    string
		AccessKey string
		SecretKey string
	}

	Payment struct {
		MidtransBaseURLProduction  string
		MidtransBaseURLDevelopment string
		MidtransServerKey          string
	}

	DistanceMatrix struct {
		DistanceMatrixAPIKey  string
		DistanceMatrixBaseURL string
	}
	Email struct {
		Domain string
		ApiKey string
	}
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func Get() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}
	return appConfig
}

func initConfig() *AppConfig {
	var config AppConfig

	config.App.Port = GetEnv("APP_PORT", "8000")
	config.App.BaseURL = GetEnv("APP_BASE_URL", "http://localhost:8000")
	config.App.ENV = GetEnv("APP_ENV", "development")

	config.Database.Host = GetEnv("DB_HOST", "localhost")
	config.Database.Port = GetEnv("DB_PORT", "3306")
	config.Database.Username = GetEnv("DB_USERNAME", "root")
	config.Database.Password = GetEnv("DB_PASSWORD", "")
	config.Database.Name = GetEnv("DB_NAME", "kost")

	config.Email.Domain = GetEnv("EMAIL", "email-domain")
	config.Email.ApiKey = GetEnv("API_KEY", "abcde545655757hjffyy")

	config.AwsS3.Bucket = GetEnv("AWS_S3_BUCKET", "kost-bucket")
	config.AwsS3.Region = GetEnv("AWS_S3_REGION", "ap-southeast-1")
	config.AwsS3.AccessKey = GetEnv("AWS_S3_ACCESS_KEY", "AKIAJXZQZQ7Z5Z5Z5Z5Z")
	config.AwsS3.SecretKey = GetEnv("AWS_S3_SECRET_KEY", "")

	config.Payment.MidtransBaseURLProduction = GetEnv("MIDTRANS_BASE_URL_PRODUCTION", "https://api.midtrans.com")
	config.Payment.MidtransBaseURLDevelopment = GetEnv("MIDTRANS_BASE_URL_DEVELOPMENT", "https://api.midtrans.com/v2/")
	config.Payment.MidtransServerKey = GetEnv("MIDTRANS_SERVER_KEY", "")

	config.DistanceMatrix.DistanceMatrixAPIKey = GetEnv("DISTANCE_MATRIX_API_KEY", "")
	config.DistanceMatrix.DistanceMatrixBaseURL = GetEnv("DISTANCE_MATRIX_BASE_URL", "https://maps.googleapis.com/maps/api/distancematrix/json")

	// Info
	fmt.Println(config.App)
	fmt.Println(config.Database)
	fmt.Println(config.AwsS3)
	fmt.Println(config.Payment)
	fmt.Println(config.DistanceMatrix)

	return &config
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		fmt.Println(value)
		return value
	}

	return fallback
}
