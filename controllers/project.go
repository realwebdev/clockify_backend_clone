package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/realwebdev/Bilal/clockify3/models"
)

func CreateProject(project models.Project, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := models.CreateProject(project, db)
		if err != nil {
			log.Print("error while creating project r", err)
			return
		}
		c.JSON(http.StatusOK, "project  Successfully created r")
	}
}

func GetProjects(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := models.GetProjects(db)
		if err != nil {
			log.Print(err)
			return
		}
		c.JSON(http.StatusOK, users)
	}
}

func UpdateProject(project_id uint, updates string, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := models.UpdateProject(project_id, updates, db)
		if err != nil {
			log.Print("error while creating project r", err)
			return
		}
		c.JSON(http.StatusOK, "project  Successfully created r")
	}
}

func DeleteProject(project_id uint, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := models.DeleteProject(project_id, db)
		if err != nil {
			log.Print("error while deleteing project r", err)
			return
		}
		c.JSON(http.StatusOK, "deleted project successfully r ")
	}
}
