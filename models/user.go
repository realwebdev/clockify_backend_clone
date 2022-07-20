package models

type User struct {
	Projects []Project  `gorm:"foreignKey:User"`
	Activity []Activity `gorm:"foreignKey:User"`

	user_id  uint `gorm:"primaryKey"`
	username string
	email    string `gorm:"typevarchar(100);unique_index"`
	password string
}
