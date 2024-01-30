package controllers

import (
	"go_tutorials/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetTasks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var tasks []models.Task
	db.Find(&tasks)
	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

func GetTask(c *gin.Context) {

	var task = map[string]interface{}{
		"id": 1,
	}
	c.JSON(http.StatusOK, gin.H{"data": task})
}
