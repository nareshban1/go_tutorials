package controllers

import (
	"go_tutorials/database"
	"go_tutorials/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// create a struct to read input from request body
type CreateTaskInput struct {
	TaskName        string `json:"taskName" binding:"required"`
	TaskDescription string `json:"taskDescription" binding:"required"`
	TaskStatus      bool   `json:"taskStatus" `
}

// api to get all tasks from database
func GetTasks(c *gin.Context) {
	// create a slice of tasks
	var tasks []models.Task
	// get all tasks from database
	database.Database.Find(&tasks)
	// return response
	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

func CreateTask(c *gin.Context) {
	// create a new task
	var newTaskData CreateTaskInput
	// validate input
	if err := c.ShouldBindJSON(&newTaskData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// create task
	newTask := models.Task{TaskName: newTaskData.TaskName, TaskDescription: newTaskData.TaskDescription, TaskStatus: newTaskData.TaskStatus}
	database.Database.Create(&newTask)
	// return response
	c.JSON(http.StatusOK, gin.H{"data": newTask})
}

func GetTask(c *gin.Context) {
	// get task by id
	var task models.Task
	// check if task exists in database or not by id provided in url param and return error if not exists
	if err := database.Database.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task not found!"})
		return
	}
	// return response if task exists
	c.JSON(http.StatusOK, gin.H{"data": task})
}

func UpdateTask(c *gin.Context) {
	// get task by id  and return error if not exists
	var task models.Task
	if err := database.Database.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task not found!"})
		return
	}

	// update task
	var updateTaskData CreateTaskInput
	// validate input
	if err := c.ShouldBindJSON(&updateTaskData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// update task in database and return response if task exists and updated successfully
	database.Database.Model(&task).Updates(updateTaskData)
	c.JSON(http.StatusOK, gin.H{"data": task})
}

func DeleteTask(c *gin.Context) {
	var task models.Task
	// check if task exists in database or not by id provided in url param and return error if not exists
	if err := database.Database.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Task not found!"})
		return
	}
	// delete task from database and return response if task exists and deleted successfully
	database.Database.Delete(&task)
	c.JSON(http.StatusOK, gin.H{"data": "Deleted Successfully"})
}
