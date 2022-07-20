package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //Gorm postgres dialect interface
	"github.com/joho/godotenv"
	"github.com/realwebdev/Bilal/clockify3/models"
)

func ConnectDB() *gorm.DB {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbname := os.Getenv("NAME")
	dbpassword := os.Getenv("PASSWORD")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbname, dbpassword, dbPort)
	var errr error
	var db *gorm.DB
	// Openning connection to database
	db, errr = gorm.Open(dialect, dbURI)

	if errr != nil {
		log.Fatal(errr)
	} else {
		fmt.Println("Connected to database successfully")
	}

	defer db.Close()
	var (
		user1 = &models.User{user_id: 1, username: "haseeb", email: "haseeb@shalan@sasketchwan.com", password: "password"}
	)
	// Loading enviroment variables

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Project{})
	db.AutoMigrate(&models.Activity{})

	db.Create(user1)

	return db

}
