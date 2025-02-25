package database

import (
	"github.com/digitaldog444/mono-monero/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeDatabase(dsn string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error while initializing database connection.")
	}

	db.AutoMigrate(&models.Keys{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.UserSettings{})
	db.AutoMigrate(&models.Wallet{})

	return db
}