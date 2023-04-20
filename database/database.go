package database

import (
	"fmt"
	"go-simple-products-api/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error

	var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		"root",
		"",
		"localhost",
		"3306",
		"echo_sample",
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("error: ", err)
	}

	DB.AutoMigrate(&models.Product{})

	log.Println("connected to the database")
}
