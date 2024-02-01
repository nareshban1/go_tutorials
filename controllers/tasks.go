package controllers

import (
	"fmt"
	"go_tutorials/database"
	"go_tutorials/models"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

// create a struct to read input from request body
type CreateTaskInput struct {
	TaskName        string `json:"taskName" binding:"required"`
	TaskDescription string `json:"taskDescription" binding:"required"`
	TaskStatus      bool   `json:"taskStatus" `
	UserID          *uint  `json:"userId"`
}

type PaginatedTasks struct {
	NextPage    int `json:"nextPage"`
	TotalPage   int `json:"totalPage"`
	CurrentPage int `json:"currentPage"`
}

// api to get all tasks from database
func GetTasks(c *gin.Context) {
	// create a slice of tasks
	var tasks []models.Task
	// get all tasks from database

	result := database.Database.Preload("User").Find(&tasks)

	// return response
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func GetPaginatedTask(c *gin.Context) {
	var tasks []models.Task
	perPageCount := 10
	page := 1
	offset := (page - 1) * int(perPageCount)
	var totalRows int64
	database.Database.Model(&models.Task{}).Count(&totalRows)
	database.Database.Limit(int(perPageCount)).Offset(offset).Find(&tasks)
	totalPages := math.Ceil(float64(totalRows) / float64(perPageCount))
	c.JSON(http.StatusOK, gin.H{"data": tasks, "pagination": PaginatedTasks{
		NextPage:    page + 1,
		CurrentPage: page,
		TotalPage:   int(totalPages),
	}})

}

func CreateTask(c *gin.Context) {
	// create a new task
	var newTaskData CreateTaskInput
	// validate input
	if err := c.ShouldBindJSON(&newTaskData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// create task  user can be assigned to a task if user id is sent
	newTask := models.Task{TaskName: newTaskData.TaskName, TaskDescription: newTaskData.TaskDescription, TaskStatus: newTaskData.TaskStatus, UserID: newTaskData.UserID}
	database.Database.Create(&newTask)
	if newTaskData.UserID != nil {
		database.Database.Preload("User").First(&newTask)
	}
	// return response
	c.JSON(http.StatusOK, gin.H{"data": newTask})
}

func GetTask(c *gin.Context) {
	// get task by id
	var task models.Task
	// check if task exists in database or not by id provided in url param and return error if not exists
	if err := database.Database.Where("id = ?", c.Param("id")).Preload("User").First(&task).Error; err != nil {
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

	updateTask := models.Task{TaskName: updateTaskData.TaskName, TaskDescription: updateTaskData.TaskDescription, TaskStatus: updateTaskData.TaskStatus, UserID: updateTaskData.UserID}
	// update task and also set user to null of null is passed from the api and dont update the primary id
	if err := database.Database.Model(&task).Select("*").Omit("ID").Updates(updateTask).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(updateTaskData)
	database.Database.Preload("User").First(&task, task.ID)
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
	if err := database.Database.Delete(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "Deleted Successfully"})
}
