package models

import (
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
	starttime := time.Now()
	activity.Start_time = starttime
	if err := db.Create(&activity).Error; err != nil {
		return err
	}
	return nil
}

func EndActivity(activity_id uint, db *gorm.DB) (time.Duration, error) {
	activity := Activity{}
	if err := db.Table("activities").Where("ID = ?", activity_id).First(&activity).Error; err != nil {
		return 0, err
	}

	starttime := activity.Start_time
	endTime := time.Now()
	difference := endTime.Sub(starttime)
	if err := db.Table("activities").Where("ID = ?", activity_id).Updates(map[string]interface{}{"total_time": difference, "end_time": endTime}).Error; err != nil {
		return 0, err
	}
	return difference, nil
}

func DeleteActivity(activityid uint, db *gorm.DB) error {
	if err := db.Table("activities").Where("ID = ?", activityid).First(&Activity{}).Error; err != nil {
		return err
	}

	if err := db.Table("activities").Where("ID=?", activityid).Delete(&Activity{}).Error; err != nil {
		return err
	}
	return nil
}

func UpdateActivity(activity_id uint, new_name string, db *gorm.DB) error {
	if err := db.Table("activities").Where("ID = ?", activity_id).First(&Activity{}).Error; err != nil {
		return err
	}

	if err := db.Table("activities").Where("ID = ?", activity_id).Updates(map[string]interface{}{"activity_name": new_name}).Error; err != nil {
		return err
	}
	return nil
}
