package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var GlobalDB *gorm.DB

func DatabaseConnection() (err error) {
	dbHost := os.Getenv("DB_HOST")
	dbUsername := os.Getenv("DB_USERNAME")
	dbDatabase := os.Getenv("DB_DATABASE")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")

	fmt.Println("DB_HOST:", dbHost)
	fmt.Println("DB_USERNAME:", dbUsername)
	fmt.Println("DB_DATABASE:", dbDatabase)
	fmt.Println("DB_PASSWORD:", dbPassword)
	fmt.Println("DB_PORT:", dbPort)
	// dbHost := "localhost"
	// dbUsername := "postgres"
	// dbDatabase := "Book_Depo_DB"
	// dbPassword := "014500"
	// dbPort := "5432"

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		dbUsername,
		dbPassword,
		dbHost,
		dbPort,
		dbDatabase)

	GlobalDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	return err
}
