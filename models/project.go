package models

type Project struct {
	Activity []Activity `gorm:"foreignKey:Project"`

	Project_id   uint `gorm:"primaryKey"`
	User_id      uint
	Project_name string
}
