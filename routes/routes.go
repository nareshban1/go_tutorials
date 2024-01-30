package routes

import (
	"go_tutorials/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.GET("/tasks", controllers.GetTasks)
	r.GET("/task/:id", controllers.GetTask)
	r.POST("/create-task", controllers.CreateTask)
	r.PATCH("/update-task/:id", controllers.UpdateTask)
	r.DELETE("/delete-task/:id", controllers.DeleteTask)

	return r
}
