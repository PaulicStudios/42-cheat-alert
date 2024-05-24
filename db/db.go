package db

import (
	"log"
	"os"

	"github.com/PaulicStudios/42-cheat-alert/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("MYSQL_DSN")

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
	log.Println("Connected to database")
	MigrateDB()
}

func MigrateDB() {
	db.AutoMigrate(&models.User{})
	log.Println("Migrated database")
}
