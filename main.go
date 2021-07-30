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

	// models.SignUp(models.User{Username: "zaseeb", Email: "kkskitchawn@gmail.com3", Password: "mypass2"}, db)
	//models.SignIn("kkskitchawn@gmail.com", "mypass2", db)
	//models.UserDeletion(3, db)

	// models.CreateProject(models.Project{Project_name: "MYFirstProj3"}, db)
	//models.DeleteProject(2, db)
	//models.UpdateProject("newname1234", "newname123455", db)

	// models.CreateActivity(models.Activity{Activity_name: "MYFirsActivity4", Start_time: time.Now(), End_time: time.Now().Add(time.Minute * 10), Total_time: time.Now().Add(time.Minute * 10)}, db)
	// models.DeleteActivity(4, db)
	// models.UpdateActivity("MYFirstActivity", "2nd ACtivity", db)

}
