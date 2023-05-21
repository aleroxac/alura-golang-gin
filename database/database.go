package database

import (
	"fmt"
	"log"
	"os"

	"github.com/aleroxac/alura-golang-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func InitDatabaseConnection() {
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_user := os.Getenv("POSTGRES_USER")
	db_pass := os.Getenv("POSTGRES_PASSWORD")
	db_dbname := os.Getenv("POSTGRES_DB")
	db_tz := os.Getenv("DB_TZ")

	connection_string := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s", db_host, db_port, db_user, db_pass, db_dbname, db_tz)
	DB, err = gorm.Open(postgres.Open(connection_string))
	if err != nil {
		log.Fatal("Error connecting to database")
	}

	DB.AutoMigrate(&models.Skill{})
}
