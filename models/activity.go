package models

import "time"

type Activity struct {
	activity_id uint `gorm:"primaryKey"`
	project_id  uint
	user_id     uint
	start_time  time.Time
	end_time    time.Time
	total_time  time.Time
}
