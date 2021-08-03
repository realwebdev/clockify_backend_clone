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

func EndActivity(activity_id uint, endtime interface{}, db *gorm.DB) {
	if err := db.Table("activities").Where("ID = ?", activity_id).Updates(map[string]interface{}{"end_time": endtime}).Error; err != nil {
		log.Print("Error occured while updating activity end time!")
		return
	}
	log.Print("Successfully updated end of activity in DB")

}

func TotalActivitytime(activity_id uint, db *gorm.DB) {
	// activity := Activity{}
	// st := db.Where(map[string]interface{}{"ID": activity_id}).First(activity.Start_time)
	// if st != nil {
	// 	log.Print("Error occured while updating activity end time!")
	// 	return
	// }
	// et := db.Where(map[string]interface{}{"ID": activity_id}).First(activity.End_time)
	// if et != nil {
	// 	log.Print("Error occured while updating activity end time!")
	// 	return
	// }
	difference := et.Sub(st)
	tt := db.Table("activities").Where("ID = ?", activity_id).Updates(map[string]interface{}{"total_time": difference})
	if tt != nil {
		log.Print("Error occured while updating activity end time!")
		return
	}
	log.Print(difference)
}

func DeleteActivity(activityid uint, db *gorm.DB) {
	if err := db.Table("activities").Where("ID=?", activityid).Delete(Activity{}).Error; err != nil {
		log.Print("Error occured activity does not exist in DB!")
		return
	}
	log.Print("Activity successfully deleted from DB")
}

func UpdateActivity(activity_id uint, new_name interface{}, db *gorm.DB) {
	if err := db.Table("activities").Where("ID = ?", activity_id).Updates(map[string]interface{}{"activity_name": new_name}).Error; err != nil {
		log.Print("Error occured while updating activity!")
		return
	}
	log.Print("Successfully updated user activity in DB")
}
