package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/realwebdev/Bilal/clockify3/controllers"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})
	router.GET("/getusers", controllers.GetUsers(db))
	return router
}
