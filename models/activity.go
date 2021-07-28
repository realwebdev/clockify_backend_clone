package models

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

type Activity struct {
	Activity_id   uint `gorm:"primaryKey"`
	Project_id    uint
	User_id       uint
	Activity_name string
	Start_time    time.Time
	End_time      time.Time
	Total_time    time.Time
}

func CreateActivity(act Activity, db *gorm.DB) {
	if err := db.Create(act).Error; err != nil {
		log.Print("Error Occured")
	}
	log.Print("your activity has been created")
}

// func startActivityUpdate(time.Time) {

// 	//start time means save time
// }

// func endActivityUpdate(time.Time) {

// 	//end time
// 	//return total time of activity
// }

func DeleteActivity(actname string, db *gorm.DB) {
	//delete row by id or name
	if err := db.Where("Activity_name LIKE ?", actname).Delete(Project{}).Error; err != nil {
		log.Print("Activity not exist")
	}
	log.Print("Activity successfully deleted")
}

func UpdateActivity(old_name string, name_update string, db *gorm.DB) {
	// change activity name
	if err := db.Model(Project{}).Where("Activity_name = ?", old_name).Updates(Activity{Activity_name: name_update}).Error; err != nil {
		log.Print("some error in update")
	}
	log.Print("successfully updated the name of user activity")
}
