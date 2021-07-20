package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

type User struct {
	Projects []Project  `gorm:"foreignKey:User"`
	Activity []Activity `gorm:"foreignKey:User"`

	user_id  uint `gorm:"primaryKey"`
	username string
	email    string `gorm:"typevarchar(100);unique_index"`
	password string
}

type Project struct {
	Activity []Activity `gorm:"foreignKey:Project"`

	project_id   uint `gorm:"primaryKey"`
	user_id      uint
	project_name string
}

type Activity struct {
	activity_id uint `gorm:"primaryKey"`
	project_id  uint
	user_id     uint
	start_time  time.Time
	end_time    time.Time
	total_time  time.Time
}

var (
	user = &User{user_id: 1, username: "haseeb", email: "haseeb@shalan@sasketchwan.com", password: "password"}
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Loading enviroment variables
	dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbname := os.Getenv("NAME")
	dbpassword := os.Getenv("PASSWORD")

	// Database connection string
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbname, dbpassword, dbPort)

	// Openning connection to database
	db, err = gorm.Open(dialect, dbURI)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to database successfully")
	}

	defer db.Close()

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Project{})
	db.AutoMigrate(&Activity{})

	db.Create(user)
}
