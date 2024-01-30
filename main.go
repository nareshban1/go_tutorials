package main

import (
	"go_tutorials/database"
	"go_tutorials/models"
	"go_tutorials/routes"
)

func main() {
	db := database.SetupDB()
	db.AutoMigrate(&models.Task{})
	r := routes.SetupRoutes(db)
	r.Run()

}
