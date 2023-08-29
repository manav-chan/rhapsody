package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/manav-chan/rhapsody/models"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file") // if Fatal error that is program stops executing
	}
	dsn := os.Getenv("DSN")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to database") // stop execution of program
	} else {
		log.Println("Connected to database")
	}
	DB = database

	// create table in database
	DB.AutoMigrate(
		&models.User{},
	)
}
