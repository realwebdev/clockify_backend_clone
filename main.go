package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/realwebdev/Bilal/clockify3/conf"
	"github.com/realwebdev/Bilal/clockify3/database"
)

func main() {
	conf.LoadEnvVar()
	db := database.ConnectDB()
	defer db.Close()
	// database.AutoMigrate(models.Activity{}, db)
	// database.AutoMigrate(models.User{}, db)
	// database.AutoMigrate(models.Project{}, db)

	// models.SignUp(models.User{Username: "zaseeb", Email: "kkskitchawn@gmail.com2", Password: "mypass2"}, db)
	// models.SignIn("kkskitchawn@gmail.com1", "mypass2", db)
	// models.UserDeletion(1, db)

	// models.CreateProject(models.Project{UserID: 3, Project_name: "MYFirstProj1"}, db)
	// models.DeleteProject(2, db)
	// models.UpdateProject(1, "newnameproject2", db)

	// models.CreateStartActivity(models.Activity{Activity_name: "MYFirsActivity", ProjectID: 1, UserID: 1, Start_time: time.Now()}, db)
	// models.EndActivity(1, time.Now(), db)
	// models.TotalActivitytime(1, time.Now(), time.Now(), db)
	// models.DeleteActivity(4, db)
	// models.UpdateActivity(1, "MYFirstActivityChanged", db)

}
