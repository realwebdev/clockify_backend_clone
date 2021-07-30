package models

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

type Activity struct {
	ID            uint
	ProjectID     uint
	UserID        uint
	Activity_name string
	Start_time    time.Time
	End_time      time.Time
	Total_time    time.Time
}

func CreateActivity(activity Activity, db *gorm.DB) {
	if err := db.Create(&activity).Error; err != nil {
		log.Print("Error Occured while creating the activity in DB!")

		return
	}
	log.Print("your activity has been created")
}

// func startActivityUpdate(time.Time) {

// 	//start time
// }

// func endActivityUpdate(time.Time) {

// 	//end time
// 	//return total time of activity
// }

func DeleteActivity(activityid uint, db *gorm.DB) {
	if err := db.Where("ID LIKE ?", activityid).Delete(Activity{}).Error; err != nil {
		log.Print("Error occured activity does not exist in DB!")

		return
	}
	log.Print("Activity successfully deleted from DB")
}

func UpdateActivity(old_name string, name_update string, db *gorm.DB) {
	if err := db.Model(Project{}).Where("Activity_name = ?", old_name).Updates(Activity{Activity_name: name_update}).Error; err != nil {
		log.Print("Error occured while updating activity!")

		return
	}
	log.Print("Successfully updated user activity in DB")
}
