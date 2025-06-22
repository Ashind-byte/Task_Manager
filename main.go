package main

import (
	"TodoApp/database"
	"TodoApp/tasks"
	"fmt"
	"log"
)

func main() {

	taskdb := database.NewTaskDb()
	taskdb.LoadDb()
	for {
		var choice int = 0
		fmt.Println("Welcome to TodoApp")
		fmt.Println("1. Create Task 2.Delete Task  3.Update Task 4.View Task ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			newTask, err := tasks.AddTask()
			if err != nil {
				log.Fatal(err)
				break
			}
			err = taskdb.AddTask(newTask)
			if err != nil {
				log.Fatal(err)
				break
			}
			break
		case 2:
			var deleteId uint = 0
			fmt.Println("Enter Id of the task to delete")
			fmt.Scan(&deleteId)
			err := taskdb.DeleteTask(deleteId)
			if err != nil {
				log.Fatal(err)
				break
			}
			break
		case 3:

			var updateId uint = 0
			fmt.Println("Enter Id of the task to update")
			fmt.Scan(&updateId)
			err := taskdb.UpdateTask(updateId)
			if err != nil {
				log.Fatal(err)
				break
			}
			break
		case 4:
			fmt.Println("Lisiting all taks...")
			err := taskdb.ListTask()
			if err != nil {
				log.Fatal(err)
				break
			}
			break
		default:
			fmt.Println("Invalid choice")
			break

		}

	}
}
