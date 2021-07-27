package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //Gorm postgres dialect interface
)

func ConnectDB() *gorm.DB {

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

	return db

}
