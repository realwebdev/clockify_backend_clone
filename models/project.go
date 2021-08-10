package models

import (
	"log"

	"github.com/jinzhu/gorm"
)

type Project struct {
	Activitys    []Activity //`gorm:"ForeignKey:ProjectID"`
	ID           uint
	UserID       uint   `json:"userid"`
	Project_name string `json:"project_name"`
}

func CreateProject(project Project, db *gorm.DB) error {
	if err := db.Create(&project).Error; err != nil {
		log.Print("Error Occured while creating project!")
		return err
	}
	log.Print("project has been created")
	return nil
}

func GetProjects(db *gorm.DB) (projects []Project, err error) {
	if err := db.Find(&projects).Error; err != nil {
		log.Print("error occured while getting user!")
		return nil, err
	}
	log.Print("These are users in database", projects)
	return projects, nil
}

func UpdateProject(project_id uint, updates string, db *gorm.DB) error {
	if err := db.Table("projects").Where("ID = ?", project_id).Updates(map[string]interface{}{"project_name": updates}).Error; err != nil {
		log.Print("Error occured while updating the project!")
		return err
	}
	log.Print("Project successfully updated ")
	return nil
}

func DeleteProject(project_id uint, db *gorm.DB) error {
	if err := db.Table("projects").Where("ID = ?", project_id).Delete(Project{}).Error; err != nil {
		log.Print("Error occured while deleting Project!")
		return err
	}
	log.Print("Project successfully deleted")
	return nil
}
