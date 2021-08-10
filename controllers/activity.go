package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/realwebdev/Bilal/clockify3/models"
)

func CreateStartActivity(activity models.Activity, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := models.CreateStartActivity(activity, db)
		if err != nil {
			log.Print("error while creating project r", err)
			return
		}
		c.JSON(http.StatusOK, "activity has been created Successfully created r")
	}
}

func EndActivity(activity_id uint, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := models.EndActivity(activity_id, db)
		if err != nil {
			log.Print("error in ending the activity time r", err)
			return
		}
		c.JSON(http.StatusOK, "activity ended r")
	}

}

func DeleteActivity(activityid uint, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := models.DeleteActivity(activityid, db)
		if err != nil {
			log.Print("error while deleting activity r", err)
			return
		}
		c.JSON(http.StatusOK, "activity has been deleted successfully r")
	}
}

func UpdateActivity(activity_id uint, new_name string, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := models.UpdateActivity(activity_id, new_name, db)
		if err != nil {
			log.Print("error while updating activity r", err)
			return
		}
		c.JSON(http.StatusOK, "activity has been updated Successfully r")
	}
}
