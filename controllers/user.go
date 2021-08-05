package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/realwebdev/Bilal/clockify3/models"
)

func GetUsers(c *gin.Context, db *gorm.DB) {
	users := models.User{}
	if err := models.GetUsers(db).Error; err != nil {
		log.Print("error in getting user")
		return
	}
	c.JSON(http.StatusOK, users)
}
