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

func CreateProject(project Project, db *gorm.DB) {
	if err := db.Create(&project).Error; err != nil {
		log.Print("Error Occured while creating project!")
		return
	}
	log.Print("project has been created")
}

func UpdateProject(project_id uint, updates string, db *gorm.DB) {
	if err := db.Table("projects").Where("ID = ?", project_id).Updates(map[string]interface{}{"project_name": updates}).Error; err != nil {
		log.Print("Error occured while updating the project!")
		return
	}
	log.Print("Project successfully updated ")
}

func DeleteProject(project_id uint, db *gorm.DB) {
	if err := db.Table("projects").Where("ID = ?", project_id).Delete(Project{}).Error; err != nil {
		log.Print("Error occured while deleting Project!")
		return
	}
	log.Print("Project successfully deleted")
}
