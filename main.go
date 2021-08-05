package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/realwebdev/Bilal/clockify3/conf"
	"github.com/realwebdev/Bilal/clockify3/database"
	"github.com/realwebdev/Bilal/clockify3/routers"
)

func main() {
	conf.LoadEnvVar()
	db := database.ConnectDB()
	defer db.Close()

	r := routers.SetupRouter()

	r.Run(":8080")
	// database.AutoMigrate(models.Activity{}, db)
	// database.AutoMigrate(models.User{}, db)
	// database.AutoMigrate(models.Project{}, db)

	// serve := gin.Default()
	// serve.Run()

	// models.GetUser(db)

}
