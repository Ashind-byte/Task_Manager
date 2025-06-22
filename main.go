package main

import (
	"TodoApp/cmd"
	"TodoApp/database"
	"log"
)

func main() {

	taskdb := database.NewTaskDb()
	err := taskdb.LoadDb()
	if err != nil {
		log.Fatal(err)
		return
	}
	err = cmd.Menu(taskdb)
	if err != nil {
		log.Fatal(err)
	}
}
