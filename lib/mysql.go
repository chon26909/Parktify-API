package lib

import (
	"context"
	"fmt"
	"parktify/models"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqlLogger struct {
	logger.Interface
}

func NewMySqlConnection() *gorm.DB {
	dsn := "parktify:1234@tcp(192.168.1.152:3306)/parktify?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: &SqlLogger{},
		DryRun: false,
	})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{}, &models.Location{})

	return db
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	stmt, _ := fc()
	fmt.Printf("\nSQL => %v\n", stmt)
}
