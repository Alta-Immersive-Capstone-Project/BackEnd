package utils

import (
	"kost/configs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlGorm(config *configs.AppConfig) *gorm.DB {

	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
	// 	config.Database.Username,
	// 	config.Database.Password,
	// 	config.Database.Host,
	// 	config.Database.Port,
	// 	config.Database.Name,
	// )
	dsn := "root@tcp(localhost:3306)/Sewa?charset=utf8mb4&parseTime=true&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return db
}
