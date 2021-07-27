package models

type User struct {
	Projects []Project  `gorm:"foreignKey:User"`
	Activity []Activity `gorm:"foreignKey:User"`

	User_id  uint `gorm:"primaryKey"`
	Username string
	Email    string `gorm:"typevarchar(100);unique_index"`
	Password string
}
