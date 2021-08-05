package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/realwebdev/Bilal/clockify3/controllers"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})
	user := controllers.GetUsers(db) //error for controllers.GetUsers(db) (no value) used as valuecompilerTooManyValues
	// too few arguments in call to controllers.GetUserscompilerWrongArgCount
	router.GET("/users", user)

	return router
}
