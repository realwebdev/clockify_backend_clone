package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/realwebdev/Bilal/clockify3/datastore"
	"github.com/realwebdev/Bilal/clockify3/models"
)

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := datastore.GetUsers()
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Error occured while retrieving user list",
				"error":   err.Error()})
			return
		}

		c.JSON(http.StatusOK, users)
	}
}

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := models.User{}
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "error occured while binding data",
				"error":   err.Error()})
			return
		}

		if err := datastore.CreateUser(user); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Error while registering user",
				"error":   err.Error()})
			return
		}

		c.JSON(http.StatusOK, fmt.Sprintf("User created %v", user.Username))
	}
}

func AuthenticateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.PostForm("email")
		pass := c.PostForm("pass")

		updates := make(map[string]interface{})
		updates["email"] = email
		updates["email"] = pass
		username, err := datastore.AuthnticateUser(updates)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "user does not exist",
				"error":   err.Error()})
			return
		}

		c.JSON(http.StatusOK, fmt.Sprintf("signedin %v", username))
	}
}

func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		uintt, _ := strconv.ParseUint(c.PostForm("id"), 10, 32)
		id := uint(uintt)

		if err := datastore.DeleteUser(id); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "user does not exist",
				"error":   err.Error()})
			return
		}

		c.JSON(http.StatusOK, "user deleted")
	}
}
