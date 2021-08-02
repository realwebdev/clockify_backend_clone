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

func CreateStartActivity(activity Activity, db *gorm.DB) {
	if err := db.Create(&activity).Error; err != nil {
		log.Print("Error Occured while Create the activity in DB!")
		return
	}
	log.Print("your activity has been created")
}

func ChangeStarttime(activity_id uint, starttime interface{}, db *gorm.DB) {
	if err := db.Table("Activity").Where("ID = ?", activity_id).Updates(map[string]interface{}{"Start_time": starttime}).Error; err != nil {
		log.Print("Error occured while updating activity start time!")
		return
	}
	log.Print("Successfully updated start activity in DB")
}

func ChangeEndtime(activity_id uint, endtime interface{}, db *gorm.DB) {
	if err := db.Table("Activity").Where("ID = ?", activity_id).Updates(map[string]interface{}{"End_time": endtime}).Error; err != nil {
		log.Print("Error occured while updating activity end time!")
		return
	}
	log.Print("Successfully updated end of activity in DB")

}

func TotaltimeActivity(activity_id uint, totaltime interface{}, db *gorm.DB) {
	if err := db.Table("Activity").Where("ID = ?", activity_id).First(&totaltime).Error; err != nil {
		log.Print("Error occured while updating activity end time!")
		return
	}
	log.Print(totaltime)
}

func DeleteActivity(activityid uint, db *gorm.DB) {
	if err := db.Table("Activity").Where("ID=?", activityid).Delete(Activity{}).Error; err != nil {
		log.Print("Error occured activity does not exist in DB!")
		return
	}
	log.Print("Activity successfully deleted from DB")
}

func UpdateActivity(activity_id uint, new_name interface{}, db *gorm.DB) {
	if err := db.Table("Activity").Where("ID = ?", activity_id).Updates(map[string]interface{}{"Activity_name": new_name}).Error; err != nil {
		log.Print("Error occured while updating activity!")
		return
	}
	log.Print("Successfully updated user activity in DB")
}
