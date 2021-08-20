package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func AutoMigrate(modelObject interface{}, db *gorm.DB) error {
	if err := db.DropTableIfExists(modelObject).Error; err != nil {
		log.Print("Error occured While Deleting Previous Table")
		return err
	}
	log.Print("Previous Table drop successfully")

	if err := db.AutoMigrate(modelObject).Error; err != nil {
		return err
	}
	log.Print("models created successfully")

	return nil
}

func ConnectDB() *gorm.DB {
	dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbname := os.Getenv("NAME")
	dbpassword := os.Getenv("PASSWORD")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbname, dbpassword, dbPort)
	db, err := gorm.Open(dialect, dbURI)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to Database!")

	return db
}
