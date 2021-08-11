package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/realwebdev/Bilal/clockify3/models"
)

func GetUsers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := models.GetUsers(db)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Error occured while retrieving user list",
				"error":   err.Error()})
			return
		}
		c.JSON(http.StatusOK, users)
	}
}

func SignUp(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := models.User{}
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "error occured while binding data",
				"error":   err.Error()})
			return
		}
		if err := models.SignUp(user, db); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Error while registering user",
				"error":   err.Error()})
			return
		}
		c.JSON(http.StatusOK, fmt.Sprintf("User created %v", user.Username))
	}
}

func SignIn(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.PostForm("email")
		pass := c.PostForm("pass")
		username, err := models.SignIn(email, pass, db)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "user does not exist",
				"error":   err.Error()})
			return
		}
		c.JSON(http.StatusOK, fmt.Sprintf("signedin %v", username))
	}
}

func UserDeletion(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		uintt, _ := strconv.ParseUint(c.PostForm("id"), 10, 32)
		id := uint(uintt)
		if err := models.UserDeletion(id, db); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "user does not exist",
				"error":   err.Error()})
			return
		}
		c.JSON(http.StatusOK, "user deleted")
	}
}
