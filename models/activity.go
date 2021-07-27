package models

import "time"

type Activity struct {
	Activity_id uint `gorm:"primaryKey"`
	Project_id  uint
	User_id     uint
	Start_time  time.Time
	End_time    time.Time
	Total_time  time.Time
}
