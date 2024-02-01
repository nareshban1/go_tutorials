package main

import (
	"go_tutorials/database"
	"go_tutorials/routes"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db := database.ConnectDatabase()
	r := routes.SetupRoutes(db)
	err := r.Run()
	if err != nil {
		panic("Something went wrong")
	}

}
