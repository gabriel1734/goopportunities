package config

import (
	"os"

	"github.com/gabriel1734/goopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Info("Database file does not exist, creating it")
		err = os.MkdirAll("./db", os.ModePerm)

		if err != nil {
			logger.Errorf("Failed to create database file: %v", err)
			return nil, err
		}

		file, err := os.Create(dbPath)

		if err != nil {
			logger.Errorf("Failed to create database file: %v", err)
			return nil, err
		}

		file.Close()
	}

	db, err := gorm.Open(sqlite.Open("./db/main.db"), &gorm.Config{})

	if err != nil {
		logger.Errorf("Failed to connect to database: %v", err)
		return nil, err
	}

	logger.Info("Connected to database")

	err = db.AutoMigrate(&schemas.Opening{})

	if err != nil {
		logger.Errorf("Failed to migrate schema: %v", err)
		return nil, err
	}

	return db, nil
}
