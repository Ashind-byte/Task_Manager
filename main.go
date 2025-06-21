package main

import (
	"TodoApp/Database"
	"TodoApp/Tasks"
	"fmt"
)

func main() {

	taskdb := Database.TaskDb{}
	taskdb.LoadDb()
	for {
		var choice int = 0
		fmt.Println("Welcome to TodoApp")
		fmt.Println("1. Create Task 2.Delete Task  3.Update Task 4.View Task ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			newTask := Tasks.AddTask()
			taskdb.AddTask(newTask)
			break
		case 2:
			var deleteId uint = 0
			fmt.Println("Enter Id of the task to delete")
			fmt.Scan(&deleteId)
			taskdb.DeleteTask(deleteId)
			break
		case 3:

			var updateId uint = 0
			fmt.Println("Enter Id of the task to update")
			fmt.Scan(&updateId)
			taskdb.UpdateTask(updateId)
			break
		case 4:
			fmt.Println("Lisiting all taks...")
			taskdb.ListTask()
			break
		default:
			fmt.Println("Invalid choice")
			break

		}

	}
}
