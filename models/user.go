package models

import (
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

func SignUp(user User, db *gorm.DB) error {
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func GetUsers(db *gorm.DB) (users []User, err error) {
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func SignIn(email, pass string, db *gorm.DB) (string, error) {
	user := User{}
	err := db.Table("users").Where(map[string]interface{}{"email": email, "password": pass}).First(&user).Error
	if err != nil {
		return "error occured", err
	}
	return user.Username, nil
}

func UserDeletion(user_id uint, db *gorm.DB) error {
	if err := db.Table("users").Where("id = ?", user_id).First(&User{}).Error; err != nil {
		return err
	}

	if err := db.Table("users").Where("id = ?", user_id).Delete(&User{}).Error; err != nil {
		return err
	}
	return nil
}
