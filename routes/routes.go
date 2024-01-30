package routes

import (
	"go_tutorials/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	//tasks
	r.GET("/tasks", controllers.GetTasks)
	r.GET("/task/:id", controllers.GetTask)
	r.POST("/create-task", controllers.CreateTask)
	r.PATCH("/update-task/:id", controllers.UpdateTask)
	r.DELETE("/delete-task/:id", controllers.DeleteTask)
	//users
	r.GET("/users", controllers.GetUsers)
	r.GET("/user/:id", controllers.GetUser)
	r.POST("/create-user", controllers.CreateUser)
	r.PATCH("/update-user/:id", controllers.UpdateUser)
	r.DELETE("/delete-user/:id", controllers.DeleteUser)

	return r
}
