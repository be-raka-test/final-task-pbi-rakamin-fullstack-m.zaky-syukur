package config

import (
	"btpn-go/app/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dbUser := "root"       // MySQL root user
	dbPass := "root"       // MySQL root password
	dbHost := "localhost"  // MySQL host
	dbPort := "3307"       // MySQL port
	dbName := "btpn_go_db" // MySQL database name

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}

	// Migrasi otomatis untuk membuat tabel
	database.AutoMigrate(&models.User{}, &models.Photo{})

	DB = database
}
