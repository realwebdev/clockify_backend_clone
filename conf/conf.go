package conf

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVar() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading env file", err)
	}
}
