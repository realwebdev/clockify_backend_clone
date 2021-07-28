package models

import "time"

type Activity struct {
	Activity_id   uint `gorm:"primaryKey"`
	Project_id    uint
	User_id       uint
	Activity_name string
	Start_time    time.Time
	End_time      time.Time
	Total_time    time.Time
}

func CreateActivity() {

	//give activity name // ID //
}

func startActivity() {

	//start time
}

func endActivity() {

	//end time
	//return total time of activity
}

func DeleteActivity() {

	//delete row by id or name
}

func UpdateActivity() {

	// change start and end time + activity name

}
