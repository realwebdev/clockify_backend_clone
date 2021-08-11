package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/realwebdev/Bilal/clockify3/models"
)

func CreateStartActivity(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var activity models.Activity
		if err := c.BindJSON(&activity); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "error occured in databinding",
				"error":   err.Error()})
			return
		}
		if err := models.CreateStartActivity(activity, db); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "error occured while creating activity",
				"error":   err.Error()})
			return
		}
		c.JSON(http.StatusOK, fmt.Sprintf("activity %v has been created!", activity.Activity_name))
	}
}

func EndActivity(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		uintt, _ := strconv.ParseUint(c.PostForm("id"), 10, 64)
		id := uint(uintt)
		t, err := models.EndActivity(id, db)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "activity not found",
				"error":   err.Error()})
			return
		}
		c.JSON(http.StatusOK, fmt.Sprintf(`total activity time =  '%v'  `, t))
	}
}

func DeleteActivity(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		uintt, _ := strconv.ParseUint(c.PostForm("id"), 10, 32)
		id := uint(uintt)
		if err := models.DeleteActivity(id, db); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "activity not found",
				"error":   err.Error()})
			return
		}
		c.JSON(http.StatusOK, "activity has been deleted!")
	}
}

func UpdateActivity(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		uintt, _ := strconv.ParseUint(c.PostForm("id"), 10, 32)
		id := uint(uintt)
		new_name := c.PostForm("newname")
		if err := models.UpdateActivity(id, new_name, db); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "activity not found",
				"error":   err.Error()})
			return
		}
		c.JSON(http.StatusOK, fmt.Sprintf("activity name has been updated to %v  !", new_name))
	}
}
