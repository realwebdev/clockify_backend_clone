package conf

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVar() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading env file")
	}
}
