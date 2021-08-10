package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/realwebdev/Bilal/clockify3/models"
)

func GetUsers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := models.GetUsers(db)
		if err != nil {
			log.Print(err)
			return
		}
		c.JSON(http.StatusOK, users)
	}
}

func SignUp(user models.User, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := models.SignUp(user, db)
		if err != nil {
			log.Print(err)
			return
		}
		c.JSON(http.StatusOK, "User Has been Created r")
	}
}

func SignIn(email, pass string, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := models.SignIn(email, pass, db)
		if err != nil {
			log.Print("error while signIn r", err)
			return
		}
		c.JSON(http.StatusOK, "user Signed In r")
	}
}

func UserDeletion(user_id uint, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := models.UserDeletion(user_id, db)
		if err != nil {
			log.Print("error while user Deletion r", err)
			return
		}
		c.JSON(http.StatusOK, "user Deleted Successfully r")
	}
}
