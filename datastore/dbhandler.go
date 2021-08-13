package datastore

import (
	"time"

	"github.com/realwebdev/Bilal/clockify3/models"
	"gorm.io/gorm"
)

type Dbhandler interface {
	CreateUser(user models.User, db *gorm.DB) error
	GetUsers(db *gorm.DB) (users []models.User, err error)
	AuthenticateUser(usercred map[string]interface{}, db *gorm.DB) (string, error)
	DeleteUser(user_id uint, db *gorm.DB) error

	CreateProject(project models.Project, db *gorm.DB) error
	GetProjects(db *gorm.DB) (projects []models.Project, err error)
	UpdateProject(project_id uint, updates map[string]interface{}, db *gorm.DB) error
	DeleteProject(project_id uint, db *gorm.DB) (string, error)

	StartActivity(activity models.Activity, db *gorm.DB) error
	EndActivity(activity_id uint, db *gorm.DB) (time.Duration, error)
	UpdateActivity(activity_id uint, updates map[string]interface{}, db *gorm.DB) error
	DeleteActivity(activityid uint, db *gorm.DB) error
}
