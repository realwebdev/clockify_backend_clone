package models

import (
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
		return err
	}
	return nil
}

func GetProjects(db *gorm.DB) (projects []Project, err error) {
	if err := db.Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}

func UpdateProject(project_id uint, updates string, db *gorm.DB) error {
	if err := db.Table("projects").Where("ID = ?", project_id).First(&Project{}).Error; err != nil {
		return err
	}

	if err := db.Table("projects").Where("ID = ?", project_id).Updates(map[string]interface{}{"project_name": updates}).Error; err != nil {
		return err
	}
	return nil
}

func DeleteProject(project_id uint, db *gorm.DB) (string, error) {
	var project Project
	if err := db.Table("projects").Where("ID = ?", project_id).First(Project{}).Error; err != nil {
		return "error occured", err
	}

	if err := db.Table("projects").Where("ID = ?", project_id).Delete(Project{}).Error; err != nil {
		return "error occured", err
	}
	return project.Project_name, nil
}
