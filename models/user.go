package models

type User struct {
	Projects  []Project  `gorm:"foreignKey:UserID"`
	Activitys []Activity `gorm:"foreignKey:UserID"`
	ID        uint
	Username  string `json:"username"`
	Email     string `gorm:"typevarchar(100);unique_index" json:"email"`
	Password  string `json:"password"`
}
