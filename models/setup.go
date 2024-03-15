package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	// database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/go_rest_api"))
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/go_rest_api"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Product{})

	DB = database
}
