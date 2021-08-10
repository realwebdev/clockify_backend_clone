package models

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

type Activity struct {
	ID            uint
	ProjectID     uint          `json:"projectid"`
	UserID        uint          `json:"userid"`
	Activity_name string        `json:"activity_name"`
	Start_time    time.Time     `json:"start_time"`
	End_time      time.Time     `json:"end_time"`
	Total_time    time.Duration `json:"total_time"`
}

func CreateStartActivity(activity Activity, db *gorm.DB) error {
	if err := db.Create(&activity).Error; err != nil {
		log.Print("Error Occured while Create the activity in DB!")
		return err
	}
	log.Print("your activity has been created")
	return nil
}

func EndActivity(activity_id uint, db *gorm.DB) error {
	activity := Activity{}
	if err := db.Table("activities").Where("ID = ?", activity_id).First(&activity).Error; err != nil {
		log.Print("Error occured while query total activity time!")
		return err
	}

	starttime := activity.Start_time
	endTime := time.Now()
	difference := endTime.Sub(starttime)
	if err := db.Table("activities").Where("ID = ?", activity_id).Updates(map[string]interface{}{"total_time": difference, "end_time": endTime}).Error; err != nil {
		log.Print("Error occured while updating total time!")
		return err
	}
	log.Print("Successfully updated total time of activity in DB")
	log.Print(difference)
	return nil
}

func DeleteActivity(activityid uint, db *gorm.DB) error {
	if err := db.Table("activities").Where("ID=?", activityid).Delete(Activity{}).Error; err != nil {
		log.Print("Error occured activity does not exist in DB!")
		return err
	}
	log.Print("Activity successfully deleted from DB")
	return nil
}

func UpdateActivity(activity_id uint, new_name string, db *gorm.DB) error {
	if err := db.Table("activities").Where("ID = ?", activity_id).Updates(map[string]interface{}{"activity_name": new_name}).Error; err != nil {
		log.Print("Error occured while updating activity!")
		return err
	}
	log.Print("Successfully updated user activity in DB")
	return nil
}
