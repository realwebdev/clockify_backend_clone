package databinding

import (
	"github.com/realwebdev/Bilal/clockify3/conf"
	"github.com/realwebdev/Bilal/clockify3/database"
	"github.com/realwebdev/Bilal/clockify3/models"
)

func DataCreate() {

	conf.LoadEnvVar()

	db := database.ConnectDB()

	// var user1 models.User
	// user1 [] = {user_id: 1, username: "haseeb", email: "haseeb@shalan@sasketchwan.com", password: "password"}
	user1 := &models.User{User_id: 1, Username: "haseeb", Email: "haseeb@shalan@sasketchwan.com", Password: "password"}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Project{})
	db.AutoMigrate(&models.Activity{})
	db.Create(user1)

}
