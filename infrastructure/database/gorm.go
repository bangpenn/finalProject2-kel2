package database

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/novita/finalproject2-revisi/config"
	"github.com/novita/finalproject2-revisi/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	Connect()
	Migration()
}

func Connect() {
	dbConfig := config.Get().Database

	// Pastikan dbConfig.Port tidak kosong sebelum mengonversi
	if dbConfig.Port == "" {
		log.Fatal("Port is empty in the database configuration")
	}

	// Convert port from string to integer
	port, err := strconv.Atoi(dbConfig.Port)
	if err != nil {
		log.Fatal("Failed to convert port to integer: ", err)
	}

	// Logging
	log.Printf("Connecting to database with user: %s, password: %s, host: %s, port: %s\n", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port)

	// Add timeout option to DSN
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable connect_timeout=5",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Name,
		port,
	)

	// Attempt to connect with timeout
	startTime := time.Now()
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	elapsedTime := time.Since(startTime)

	if err != nil {
		log.Fatal("Failed to connect to database, ", err)
	}

	log.Printf("Connected to database in %s\n", elapsedTime)
}

func Migration() {
	db.AutoMigrate(
		&entity.User{},
		&entity.Photo{},
		&entity.Comment{},
		&entity.SocialMedia{},
	)
}

func GetInstance() *gorm.DB {
	if db == nil {
		log.Fatal("Database instance is not initialized")
	}

	return db
}
