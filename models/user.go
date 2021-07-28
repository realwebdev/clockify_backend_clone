package models

import (
	"log"

	"github.com/jinzhu/gorm"
)

type User struct {
	Projects []Project  `gorm:"foreignKey:User"`
	Activity []Activity `gorm:"foreignKey:User"`

	User_id  uint `gorm:"primaryKey; autoIncrement:true"`
	Username string
	Email    string `gorm:"typevarchar(100);unique_index"`
	Password string
}

func CreateSchema(db *gorm.DB) {
	db.AutoMigrate(User{})
	db.AutoMigrate(Project{})
	db.AutoMigrate(Activity{})
}

func SignUp(user1 User, db *gorm.DB) {
	if err := db.Create(user1).Error; err != nil {
		log.Print("wrong username or password")
	}
	log.Print("your account has been created")
}

func SignIn(email string, pass string, db *gorm.DB) {
	var user User
	if err := db.Where(&User{Email: email, Password: pass}).First(&user).Error; err != nil {
		log.Print("this user not exist")
	}
	log.Print("SgnedIn")
}

func UserDeletion(email string, db *gorm.DB) {
	if err := db.Where("Email LIKE ?", email).Delete(User{}).Error; err != nil {
		log.Print("account not exist")
	}
	log.Print("account successfully deleted")
}
