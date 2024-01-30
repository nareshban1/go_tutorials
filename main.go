package main

import (
	"go_tutorials/database"
	"go_tutorials/routes"
)

func main() {
	db := database.ConnectDatabase()
	r := routes.SetupRoutes(db)
	err := r.Run()
	if err != nil {
		panic("Something went wrong")
	}

}
