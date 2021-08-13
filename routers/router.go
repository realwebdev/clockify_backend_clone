package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/realwebdev/Bilal/clockify3/handlers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	router.GET("/getusers", handlers.GetUsers())
	router.POST("/signup", handlers.CreateUser())
	router.PUT("/signin", handlers.AuthenticateUser())
	router.DELETE("/deleteuser", handlers.DeleteUser())

	router.GET("/getproject", handlers.GetProjects())
	router.POST("/createproject", handlers.CreateProject())
	router.POST("/updateproject", handlers.UpdateProject())
	router.DELETE("/deleteproject", handlers.DeleteProject())

	router.POST("/createactivity", handlers.StartActivity())
	router.POST("/endactivity", handlers.EndActivity())
	router.POST("/updateactivity", handlers.UpdateActivity())
	router.DELETE("/deleteactivity", handlers.DeleteActivity())

	return router
}
