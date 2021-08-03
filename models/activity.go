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
	Total_time    time.Duration
}

func CreateStartActivity(activity Activity, db *gorm.DB) {
	if err := db.Create(&activity).Error; err != nil {
		log.Print("Error Occured while Create the activity in DB!")
		return
	}
	log.Print("your activity has been created")
}

func EndActivity(activity_id uint, db *gorm.DB) {
	activity := Activity{}
	if err := db.Table("activities").Where("ID = ?", activity_id).First(&activity).Error; err != nil {
		log.Print("Error occured while query total activity time!")
		return
	}

	starttime := activity.Start_time
	endTime := time.Now()
	difference := endTime.Sub(starttime)
	if err := db.Table("activities").Where("ID = ?", activity_id).Updates(map[string]interface{}{"total_time": difference, "end_time": endTime}).Error; err != nil {
		log.Print("Error occured while updating total time!")
		return
	}
	log.Print("Successfully updated total time of activity in DB")
	log.Print(difference)
}

func DeleteActivity(activityid uint, db *gorm.DB) {
	if err := db.Table("activities").Where("ID=?", activityid).Delete(Activity{}).Error; err != nil {
		log.Print("Error occured activity does not exist in DB!")
		return
	}
	log.Print("Activity successfully deleted from DB")
}

func UpdateActivity(activity_id uint, new_name string, db *gorm.DB) {
	if err := db.Table("activities").Where("ID = ?", activity_id).Updates(map[string]interface{}{"activity_name": new_name}).Error; err != nil {
		log.Print("Error occured while updating activity!")
		return
	}
	log.Print("Successfully updated user activity in DB")
}
