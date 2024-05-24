package db

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("MYSQL_DSN")

	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	println("Connected to database")
	// MigrateDB()
}

func MigrateDB() {
	db.AutoMigrate()
	println("Migrated database")
}
