package models

import (
	"log"

	"github.com/jinzhu/gorm"
)

type Project struct {
	Activity []Activity `gorm:"foreignKey:Project"`

	Project_id   uint `gorm:"primaryKey"`
	User_id      uint
	Project_name string
}

func CreateProject(proj Project, db *gorm.DB) {
	if err := db.Create(proj).Error; err != nil {
		log.Print("Error Occured")
	}
	log.Print("your project has been created")
}

func UpdateProject(old_name string, name_update string, db *gorm.DB) {
	if err := db.Model(Project{}).Where("Project_name = ?", old_name).Updates(Project{Project_name: name_update}).Error; err != nil {
		log.Print("some error in update")
	}
	log.Print("successfully updated the name of project")
}

func DeleteProject(projname string, db *gorm.DB) {
	if err := db.Where("Project_name LIKE ?", projname).Delete(User{}).Error; err != nil {
		log.Print("Project not exist")
	}
	log.Print("Project successfully deleted")
}
