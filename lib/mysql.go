package lib

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySqlConnection() *gorm.DB {
	dsn := "root:1234@tcp(159.138.254.50:3306)/parktify?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
