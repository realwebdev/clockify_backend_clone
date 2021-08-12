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
	router.POST("/signup", controllers.SignUp(db))
	router.PUT("/signin", controllers.SignIn(db))
	router.DELETE("/deleteuser", controllers.UserDeletion(db))

	router.GET("/getproject", controllers.GetProjects(db))
	router.POST("/createproject", controllers.CreateProject(db))
	router.POST("/updateproject", controllers.UpdateProject(db))
	router.DELETE("/deleteproject", controllers.DeleteProject(db))

	router.POST("/createactivity", controllers.StartActivity(db))
	router.POST("/endactivity", controllers.EndActivity(db))
	router.POST("/updateactivity", controllers.UpdateActivity(db))
	router.DELETE("/deleteactivity", controllers.DeleteActivity(db))

	return router
}
