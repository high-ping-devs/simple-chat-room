package database

import (
	"fmt"
	"log"

	"os"

	"github.com/high-ping-devs/simple-chat-room/backend/models"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var (
	host     = os.Getenv("DB_HOST")
	port     = os.Getenv("DB_PORT")
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbname   = os.Getenv("DB_NAME")
	sslmode  = os.Getenv("DB_SSLMODE")
	connStr  = fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v", host, port, user, password, dbname, sslmode)
	DB       *gorm.DB
)

func Connect() {
	var err error

	log.Println("📦 Connecting to database... ")

	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Panic("❌ Error connecting to database: ", err)
	}

	if DB, err := DB.DB(); err == nil {
		DB.SetMaxOpenConns(10)
	} else {
		log.Panic("❌ Error setting max open connections: ", err)
	}

	log.Println("📦 Database connected successfully")
}

func Migrate() {
	log.Println("📦 Migrating database... ")

	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Panic("❌ Error migrating database: ", err)
	}

	log.Println("📦 Database migrated successfully")
}
