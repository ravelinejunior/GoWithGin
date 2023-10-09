package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ravelinejunior/go_api_gin/controller"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/users", controller.ShowAllUsers)
	r.GET("/:name", controller.Greetings)
	r.GET("/users/:id", controller.GetUserById)
	r.POST("/users", controller.CreateNewUser)
	r.DELETE("/users/:id", controller.DeleteUser)
	r.Run(":8080")
}
