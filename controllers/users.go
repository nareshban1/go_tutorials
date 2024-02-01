package controllers

import (
	"go_tutorials/database"
	"go_tutorials/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// create a struct to read input from request body
type CreateUserInput struct {
	UserName string `json:"userName" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type UserResponse struct {
	UserName string
	Email    string
	ID       uint
}

// api to get all users from database
func GetUsers(c *gin.Context) {
	// create a slice of users
	var users []models.User
	// get all users from database
	if err := database.Database.Select("UserName", "Email").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	// return response
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func SignUp(c *gin.Context) {
	var signUpBody struct {
		UserName string `json:"userName" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&signUpBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(signUpBody.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user := models.User{UserName: signUpBody.UserName, Password: string(hash), Email: signUpBody.Email}
	if err := database.Database.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "User Created Successfully"})
}

func Login(c *gin.Context) {
	var loginBody struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var user models.User
	if err := c.ShouldBindJSON(&loginBody); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.Database.Where("email = ?", loginBody.Email).First(&user).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": loginBody.Email})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginBody.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_TOKEN")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})

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
	if err := database.Database.First(&user, c.Param("id")).Error; err != nil {
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
	if err := database.Database.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": true})
}
