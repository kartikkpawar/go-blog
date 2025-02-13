package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kartikkpawar/go-blog/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBConnection *gorm.DB

func ConnectDB() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading env file")
	}
	dsn := os.Getenv("DB_URL")
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
