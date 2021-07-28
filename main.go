package main

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/realwebdev/Bilal/clockify3/conf"
	"github.com/realwebdev/Bilal/clockify3/database"
	"github.com/realwebdev/Bilal/clockify3/models"
)

func main() {
	conf.LoadEnvVar()
	db := database.ConnectDB()
	defer db.Close()
	//models.CreateSchema(db)

	//models.SignUp(models.User{Username: "zaseeb", Email: "kkskitchawn@gmail.com", Password: "mypass2"}, db)
	//models.SignIn("kkskitchawn@gmail.com", "mypass2", db)
	//models.UserDeletion("kkskitchawn@gmail.com", db)

	// models.CreateProject(models.Project{Project_name: "MYFirstProj"}, db)
	//models.DeleteProject("MYFirstProj", db)
	//models.UpdateProject("newname1234", "newname123455", db)

	models.CreateActivity(models.Activity{Activity_name: "MYFirsActivity4", Start_time: time.Now(), End_time: time.Now().Add(time.Minute * 10), Total_time: time.Now().Add(time.Minute * 10)}, db)
	// models.DeleteActivity("MYFirstActivity", db)
	// models.UpdateActivity("MYFirstActivity", "2nd ACtivity", db)

}
