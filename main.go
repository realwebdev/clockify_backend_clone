package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	databinding "github.com/realwebdev/Bilal/clockify3/data_binding"
)

func main() {

	// Database connection string

	databinding.DataCreate()

}
