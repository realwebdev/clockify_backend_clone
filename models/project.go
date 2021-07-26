package models

type Project struct {
	Activity []Activity `gorm:"foreignKey:Project"`

	project_id   uint `gorm:"primaryKey"`
	user_id      uint
	project_name string
}
