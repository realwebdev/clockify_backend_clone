package conf

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type envvar struct {
	DIALECT  string
	HOST     string
	DBPORT   string
	USERNAME string
	PASSWORD string
}

func LoadEnvVar() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading env file", err)
	}

	ConnectDB()
}

//create constructors
//also contain func to connect to db
func ConnectDB() (string, string) {
	dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbname := os.Getenv("NAME")
	dbpassword := os.Getenv("PASSWORD")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbname, dbpassword, dbPort)

	return dialect, dbURI
}
