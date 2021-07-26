package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/realwebdev/Bilal/clockify3/database"
)

func main() {

	// Database connection string

	database.ConnectDB()

}
