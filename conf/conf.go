package conf

import (
	"log"

	"github.com/joho/godotenv"
)

func env_varr() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}
