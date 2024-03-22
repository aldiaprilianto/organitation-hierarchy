package config

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres" // Menggunakan driver PostgreSQL
	"gorm.io/gorm"
	"log"
	"os"
)

var db *gorm.DB

func InitializeDatabase() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		return errors.New("missing DB_USER")
	}

	dbPass := os.Getenv("DB_PASS")
	if dbPass == "" {
		return errors.New("missing DB_PASS")
	}
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		return errors.New("missing DB_HOST")
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		return errors.New("missing DB_NAME")
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		return errors.New("missing DB_PORT")
	}

	// Mengganti string koneksi untuk PostgreSQL
	dbString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPass, dbName, dbPort)

	database, err := gorm.Open(postgres.Open(dbString), &gorm.Config{})
	if database == nil {
		fmt.Errorf("database connection is not initialized")
	}

	//runMigrations(dbString)

	db = database

	return nil
}

func GetDB() (*gorm.DB, error) {
	if db == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}
	return db, nil
}

func CloseDatabase() error {
	db, err := GetDB()
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get *sql.DB object: %v", err)
	}

	err = sqlDB.Close()
	if err != nil {
		return fmt.Errorf("failed to close database connection: %v", err)
	}

	return nil
}
