package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/realwebdev/Bilal/clockify3/models"
)

func CreateProject(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		project := models.Project{}
		if err := c.BindJSON(&project); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "error occured in data binding",
				"error":   err.Error()})
			return
		}
		err := models.CreateProject(project, db)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "error occured creating project",
				"error":   err.Error()})
			return
		}
		c.JSON(http.StatusOK, fmt.Sprintf("project  created %v", project.Project_name))
	}
}

func GetProjects(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := models.GetProjects(db)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Project list not found",
				"error":   err.Error()})
			return
		}
		c.JSON(http.StatusOK, users)
	}
}

func UpdateProject(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		uintt, _ := strconv.ParseUint(c.PostForm("id"), 10, 32)
		id := uint(uintt)
		update := c.PostForm("update")
		if err := models.UpdateProject(id, update, db); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "The project does not exist in the database",
				"error":   err.Error()})
			return
		}
		c.JSON(http.StatusOK, fmt.Sprintf(`project name changed to '%v'`, update))
	}
}

func DeleteProject(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		uintt, _ := strconv.ParseUint(c.PostForm("id"), 10, 32)
		id := uint(uintt)
		deleteproject, err := models.DeleteProject(id, db)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "The project does not exist in the database",
				"error":   err.Error()})
			return
		}
		c.JSON(http.StatusOK, fmt.Sprintf("project %v  deleted !", deleteproject))
	}
}
