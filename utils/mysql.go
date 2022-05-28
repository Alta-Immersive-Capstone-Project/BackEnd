package utils

import (
	"fmt"
	"kost/configs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysqlGorm(config *configs.AppConfig) *gorm.DB {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		config.Username,
		config.Password,
		config.Address,
		config.Port,
		config.Name,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return db
}
