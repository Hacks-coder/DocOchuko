package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/root/medicalBooking/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBConn *gorm.DB

func ConnectDB() {

	godotenv.Load()

	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	data_base_name := os.Getenv("DBNAME")

	connection := user + ":" + password + "@tcp(127.0.0.1:3306)/" + data_base_name + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		panic("error in connecting to the database... biko check the connection")
	}
	log.Println("your connection has been done successfully...")

	db.AutoMigrate(
		&models.User{},
	)

	DBConn = db

}
