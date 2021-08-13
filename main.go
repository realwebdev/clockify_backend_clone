package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/realwebdev/Bilal/clockify3/conf"
	"github.com/realwebdev/Bilal/clockify3/routers"
)

var Db gorm.DB

func main() {
	conf.LoadEnvVar()

	// database.AutoMigrate(models.Activity{}, db)
	// database.AutoMigrate(models.User{}, db)
	// database.AutoMigrate(models.Project{}, db)

	// models.SignUp(models.User{Username: "haseeb", Email: "saskitchawn@gmail.com3.3", Password: "mypass2"}, db)
	// models.SignIn("saskitchawn@gmail.com1", "mypass1", db)
	// models.UserDeletion(1, db)

	// models.CreateProject(models.Project{UserID: 2, Project_name: "MYFirstProj3"}, db)
	// models.DeleteProject(2, db)
	// updates := make(map[string]interface{})
	// updates["project_name"] = "mynewname"
	// models.UpdateProject(1, updates, db)

	// models.CreateStartActivity(models.Activity{Activity_name: "MYFirsActivity2", ProjectID: 1, UserID: 1, Start_time: time.Now()}, db)
	// models.EndActivity(1, db)
	// models.DeleteActivity(4, db)
	// models.UpdateActivity(1, "MYFirstActivityChanged", db)

	r := routers.SetupRouter()
	r.Run(":8080")
}
