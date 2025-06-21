package Tasks

import (
	"bufio"
	"fmt"
	"gorm.io/gorm"
	"os"
	"time"
)

type Task struct {
	gorm.Model
	Title     string
	AddedDate time.Time
	Completed bool
}

func AddTask() *Task {
	reader := bufio.NewReader(os.Stdin)
	task := &Task{}
	fmt.Println("Enter the title of the task")
	task.Title, _ = reader.ReadString('\n')
	fmt.Println("Is the task complete? true or false")
	_, err := fmt.Scan(&task.Completed)
	if err != nil {
		return nil
	}
	return task

}
