package routes

import (
	"github.com/gin-gonic/gin"
	"server-mqtt/app/controllers"
	"server-mqtt/app/middleware"
)

func Routes(app *gin.Engine) {
	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to API",
		})
	})

	route := app.Use(middleware.CekAPIKey)
	ctrl := controllers.NewSensorController()
	route.GET("/sensor", ctrl.Index)
	route.GET("/sensor/:id", ctrl.Show)

	task := controllers.NewTaskController()
	route.POST("/task", task.Doanything)
	route.GET("/task", task.Index)
}
