package models

import (
	"log"

	"github.com/jinzhu/gorm"
)

type Project struct {
	Activitys    []Activity //`gorm:"ForeignKey:ProjectID"`
	ID           uint
	UserID       uint
	Project_name string
}

func CreateProject(proj Project, db *gorm.DB) {
	if err := db.Create(&proj).Error; err != nil {
		log.Print("Error Occured while creating project!")

		return
	}
	log.Print("project has been created")
}

func UpdateProject(old_name string, name_update string, db *gorm.DB) {
	if err := db.Model(Project{}).Where("Project_name = ?", old_name).Updates(Project{Project_name: name_update}).Error; err != nil {
		log.Print("Error occured while updating the project!")

		return
	}
	log.Print("Project successfully updated ")
}

func DeleteProject(projname string, db *gorm.DB) {
	if err := db.Where("Project_name LIKE ?", projname).Delete(Project{}).Error; err != nil {
		log.Print("Error occured while deleting Project!")

		return
	}
	log.Print("Project successfully deleted")
}
