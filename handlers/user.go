package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/realwebdev/Bilal/clockify3/auth"
	"github.com/realwebdev/Bilal/clockify3/models"
)

type TokenResponse struct {
	AccessToken  string
	RefreshToken string
}

func GetUsers(h *Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := auth.TokenValidate(c.Request); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Error in Authorizaition of JWT",
				"error":   err,
			})
			return
		}

		users, err := h.DB.GetUsers()
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Error occured while retrieving user list",
				"error":   err.Error()})
			return
		}

		c.JSON(http.StatusOK, users)
	}
}

func CreateUser(h *Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := models.User{}
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "error occured while binding user data",
				"error":   err.Error()})
			return
		}

		if err := h.DB.CreateUser(user); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Error while registering user",
				"error":   err.Error()})
			return
		}

		c.JSON(http.StatusOK, fmt.Sprintf("User created %v", user.Username))
	}
}

func AuthenticateUser(h *Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.PostForm("email")
		pass := c.PostForm("pass")

		updates := make(map[string]interface{})
		updates["email"] = email
		updates["password"] = pass
		username, err := h.DB.AuthenticateUser(updates)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "user does not exist",
				"error":   err.Error()})
			return
		}

		accessToken, err := auth.GenerateJWT(email)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Error Generating JWT",
				"error":   err.Error(),
			})
			return
		}

		refreshToken, err := auth.RefreshJWT(email)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Error Generating JWT",
				"error":   err.Error(),
			})
			return
		}

		response := TokenResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}
		c.JSON(http.StatusOK, fmt.Sprintf("Loged In user: %s", username))
		c.JSON(http.StatusOK, fmt.Sprintf("Access Token :%s", response.AccessToken))
		c.JSON(http.StatusOK, fmt.Sprintf(" Refresh Token :%s", response.RefreshToken))
	}
}

func DeleteUser(h *Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := auth.TokenValidate(c.Request); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Error in Authorizaition of JWT",
				"error":   err,
			})
			return
		}

		uintt, _ := strconv.ParseUint(c.PostForm("id"), 10, 32)
		id := uint(uintt)

		if err := h.DB.DeleteUser(id); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "user does not exist",
				"error":   err.Error()})
			return
		}

		c.JSON(http.StatusOK, "user deleted")
	}
}
