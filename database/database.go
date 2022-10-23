package database

import (
	"echo-item/model"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

var (
	DB_USERNAME string = os.Getenv("DB_USERNAME") //util.GetConfig("DB_USERNAME")
	DB_PASSWORD string = os.Getenv("DB_PASSWORD") //util.GetConfig("DB_PASSWORD")
	DB_NAME     string = os.Getenv("DB_NAME")     //util.GetConfig("DB_NAME")
	DB_HOST     string = os.Getenv("DB_HOST")     //util.GetConfig("DB_HOST")
	DB_PORT     string = os.Getenv("DB_PORT")     //util.GetConfig("DB_PORT")
)

func Connect() {
	var err error

	var dsn string = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		DB_HOST,
		DB_PORT,
		DB_USERNAME,
		DB_PASSWORD,
		DB_NAME,
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("error when connecting to the database: %s", err)
	}

	log.Println("connected to the database")

	DB.AutoMigrate(&model.Item{}, &model.Category{})
}
