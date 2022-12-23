package config

import (
	"fmt"
	"log"
	"os"

	"github.com/ashfaqrobin/golang-mysql-bookstore-api/pkg/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	database := os.Getenv("DB_DATABASE")

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)

	d, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic("Couldn't connect to database")
	}

	db = d

	fmt.Println("Database connected")

	// Database Migration
	err = models.MigrateBooks(d)

	if err != nil {
		log.Fatal("Couldn't migrate books")
	}

	fmt.Println("Migration completed")
}

func GetDB() *gorm.DB {
	return db
}
