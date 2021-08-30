package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/realwebdev/Bilal/clockify3/handlers"
)

func SetupRouter(handler *handlers.Handler) *gin.Engine {
	router := gin.Default()
	router.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	router.GET("/getusers", handlers.GetUsers(handler))
	router.POST("/signup", handlers.CreateUser(handler))
	router.PUT("/signin", handlers.AuthenticateUser(handler))
	router.DELETE("/deleteuser", handlers.DeleteUser(handler))

	router.GET("/getproject", handlers.GetProjects(handler))
	router.POST("/createproject", handlers.CreateProject(handler))
	router.POST("/updateproject", handlers.UpdateProject(handler))
	router.DELETE("/deleteproject", handlers.DeleteProject(handler))

	router.POST("/createactivity", handlers.StartActivity(handler))
	router.POST("/endactivity", handlers.EndActivity(handler))
	router.POST("/updateactivity", handlers.UpdateActivity(handler))
	router.DELETE("/deleteactivity", handlers.DeleteActivity(handler))

	// router.GET("/isAuthenticated", handlers.DeleteActivity(handler))

	return router
}
