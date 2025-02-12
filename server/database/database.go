package database

import (
	"log"

	"github.com/kartikkpawar/go-blog/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBConnection *gorm.DB

func ConnectDB() {
	dsn := "root:root@tcp(127.0.0.1:3306)/go_fiber_blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		panic("Database connection failed")
	}

	log.Println("DB Connection Successfull")
	db.AutoMigrate(new(model.Blog))

	DBConnection = db
}
