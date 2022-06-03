package configs

import (
	"fmt"
	"os"
	"sync"
)

type AppConfig struct {
	App struct {
		BaseURL  string
		FrontURL string
		Port     string
		ENV      string
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
		MidtransServerKey string
		MidtransStatus    string
	}

	DistanceMatrix struct {
		DistanceMatrixAPIKey  string
		DistanceMatrixBaseURL string
	}

	GCalendar struct {
		ClientID     string
		ClientSecret string
		ProjectID    string
		AuthUri      string
		AuthProvider string
		TokenUri     string
		RedirectUri  string
	}

	Email struct {
		Domain string
		ApiKey string
	}

	Frontend struct {
		Domain    string
		ResetPage string
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
	config.App.FrontURL = GetEnv("APP_FRONT_URL", "http://localhost:3000/transactions/finish")
	config.App.ENV = GetEnv("APP_ENV", "development")

	config.Database.Host = GetEnv("DB_HOST", "localhost")
	config.Database.Port = GetEnv("DB_PORT", "3306")
	config.Database.Username = GetEnv("DB_USERNAME", "root")
	config.Database.Password = GetEnv("DB_PASSWORD", "")
	config.Database.Name = GetEnv("DB_NAME", "")

	config.Email.Domain = GetEnv("EMAIL", "")
	config.Email.ApiKey = GetEnv("API_KEY", "")

	config.Frontend.Domain = GetEnv("FRONTEND_DOMAIN", "http://localhost:8001")
	config.Frontend.ResetPage = GetEnv("FRONTEND_RESET_PAGE", "/password-confirmation")
	// Info
	config.AwsS3.Bucket = GetEnv("AWS_S3_BUCKET", "")
	config.AwsS3.Region = GetEnv("AWS_S3_REGION", "")
	config.AwsS3.AccessKey = GetEnv("AWS_S3_ACCESS_KEY", "")
	config.AwsS3.SecretKey = GetEnv("AWS_S3_SECRET_KEY", "")

	config.Payment.MidtransServerKey = GetEnv("MIDTRANS_SERVER_KEY", "")
	config.Payment.MidtransStatus = GetEnv("MIDTRANS_STATUS", "")

	config.DistanceMatrix.DistanceMatrixAPIKey = GetEnv("DISTANCE_MATRIX_API_KEY", "")
	config.DistanceMatrix.DistanceMatrixBaseURL = GetEnv("DISTANCE_MATRIX_BASE_URL", "")

	config.GCalendar.ClientID = GetEnv("GCAL_OAUTH_CLIENT_ID", "350091276191-i1bcs17inok4p2ehmmmfk4s29rohag3o.apps.googleusercontent.com")
	config.GCalendar.ClientSecret = GetEnv("GCAL_OAUTH_CLIENT_SECRET", "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")
	// config.GCalendar.ClientSecret = GetEnv("GCAL_OAUTH_CLIENT_SECRET", "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")
	config.GCalendar.ProjectID = GetEnv("GCAL_PROJECT_ID", "capstone-be")
	config.GCalendar.AuthUri = GetEnv("GCAL_AUTH_URI", "https://accounts.google.com/o/oauth2/auth")
	config.GCalendar.TokenUri = GetEnv("GCAL_TOKEN_URI", "https://oauth2.googleapis.com/token")
	config.GCalendar.AuthProvider = GetEnv("GCAL_AUTH_PROVIDER", "https://www.googleapis.com/oauth2/v1/certs")
	config.GCalendar.RedirectUri = GetEnv("GCAL_REDIRECT_URI", "http://localhost:8000/auth/google/callback")
	// config.GCalendar.RedirectUri = GetEnv("GCAL_REDIRECT_URI", "https://galihgpr.cloud.okteto.net")
	// Info
	fmt.Println(config.App)
	fmt.Println(config.Database)
	fmt.Println(config.Email)
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
