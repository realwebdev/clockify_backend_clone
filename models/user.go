package models

import (
	"github.com/realwebdev/Bilal/clockify3/conf"
	"github.com/realwebdev/Bilal/clockify3/database"
)

type User struct {
	Projects []Project  `gorm:"foreignKey:User"`
	Activity []Activity `gorm:"foreignKey:User"`

	User_id  uint `gorm:"primaryKey"`
	Username string
	Email    string `gorm:"typevarchar(100);unique_index"`
	Password string
}

func CreateScema() {
	conf.LoadEnvVar()
	db := database.ConnectDB()
	defer db.Close()

	db.AutoMigrate(User{})
	db.AutoMigrate(Project{})
	db.AutoMigrate(Activity{})
}

func SignUp(user1 User) {
	// user id //user name// email // password // email check for unique user

	conf.LoadEnvVar()

	db := database.ConnectDB()
	defer db.Close()

	// var user1 models.User
	// user1 [] = {user_id: 1, username: "haseeb", email: "haseeb@shalan@sasketchwan.com", password: "password"}
	// user1 := User{User_id: 1, Username: "haseeb", Email: "haseeb@shalan@sasketchwan.com", Password: "password"}

	db.Create(user1)

	SignIn(user1.Email, user1.Password)

}

func SignIn(email string, pass string) {

	db := database.ConnectDB()
	defer db.Close()

	// Get all matched records
	db.Where(&User{Email: email, Password: pass}).Find( /*what to put here*/ )
	//check user email and password and show its profile

}

func UserDeletion() {

	// after signin or signup if id or email match user can delete by email  its whole record

}
