package controllers

import (
	"go_tutorials/database"
	"go_tutorials/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// create a struct to read input from request body
type CreateUserInput struct {
	UserName string `json:"userName" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

// api to get all users from database
func GetUsers(c *gin.Context) {
	// create a slice of users
	var users []models.User
	// get all users from database
	database.Database.Find(&users)
	// return response
	c.JSON(http.StatusOK, gin.H{"data": users})
}

// api to create a user
func CreateUser(c *gin.Context) {
	// create an empty user struct
	var input CreateUserInput
	// validate input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// create a user
	user := models.User{UserName: input.UserName, Email: input.Email}
	// save user in database
	database.Database.Create(&user)
	// return response
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// api to get a user by id
func GetUser(c *gin.Context) {
	// create an empty user struct
	var user models.User
	// check if user exists in database or not by id provided in url param and return error if not exists
	if err := database.Database.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}
	// return response if user exists
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// api to update a user by id
func UpdateUser(c *gin.Context) {
	var user models.User
	// check if user exists in database or not by id provided in url param and return error if not exists
	if err := database.Database.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}
	// update user
	var input CreateUserInput
	// validate input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// update user in database and return response if user exists and updated successfully

	database.Database.Model(&user).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// api to delete a user by id
func DeleteUser(c *gin.Context) {
	var user models.User
	// check if user exists in database or not by id provided in url param and return error if not exists
	if err := database.Database.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}
	// delete user from database and return response if user exists and deleted successfully
	database.Database.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
