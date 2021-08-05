package models

import (
	"log"

	"github.com/jinzhu/gorm"
)

type User struct {
	Projects  []Project  //`gorm:"foreignKey:UserID"`
	Activitys []Activity //`gorm:"foreignKey:UserID"`
	ID        uint
	Username  string `json:"username"`
	Email     string `gorm:"typevarchar(100);unique_index" json:"email"`
	Password  string `json:"password"`
}

func SignUp(user User, db *gorm.DB) {
	if err := db.Create(&user).Error; err != nil {
		log.Print("Error occured While creating the user in DB!")
		return
	}
	log.Print("your account has been created")
}

func GetUsers(db *gorm.DB) error {
	user := User{}
	if err := db.Find(&user).Error; err != nil {
		log.Print("error occured while getting user!")
		return err
	}
	log.Print(user)
	return nil
}

func SignIn(email, pass string, db *gorm.DB) {
	user := User{}
	if err := db.Table("users").Where(map[string]interface{}{"email": email, "password": pass}).Find(&user).Error; err != nil {
		log.Print("Error occured while SignIn")
		return
	}
	log.Print("User SignedIn")
}

func UserDeletion(user_id uint, db *gorm.DB) {
	if err := db.Table("User").Where("ID = ?", user_id).Delete(User{}).Error; err != nil {
		log.Print("Error occured while deleting user!")
		return
	}
	log.Print("user deleted")
}
