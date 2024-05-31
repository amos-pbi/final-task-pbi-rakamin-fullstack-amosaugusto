package main

import (
	"fmt"

	"github.com/amos-pbi/final-task-pbi-rakamin-fullstack-amosaugusto/database"
	"github.com/amos-pbi/final-task-pbi-rakamin-fullstack-amosaugusto/router"
)

func main() {
	database.ConnectDatabase()

	r := router.SetupRouter()
	r.Run("localhost:8080")
	if err := r.Run(); err != nil {
		fmt.Println("Error starting server: ", err)
		return
	}
}
