package Database

import (
	"TodoApp/Tasks"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type TaskDb struct {
	db *gorm.DB
}

func (t *TaskDb) LoadDb() {
	var err error
	t.db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	if err := t.db.AutoMigrate(&Tasks.Task{}); err != nil {
		log.Fatal(err)
	}

}

func (t *TaskDb) AddTask(task *Tasks.Task) {
	result := t.db.Create(task)

	if result.Error != nil {
		log.Fatal(result.Error)
		return
	}
	fmt.Println("Added task successfully")
}

func (t *TaskDb) UpdateTask(taskId uint) {
	t.db.Model(&Tasks.Task{}).Where("id = ?", taskId).Updates(Tasks.Task{Completed: true})
	fmt.Println("Updated task successfully")
}

func (t *TaskDb) DeleteTask(taskId uint) {
	t.db.Model(&Tasks.Task{}).Where("id = ?", taskId).Delete(&Tasks.Task{})
	fmt.Println("Deleted task successfully")
}

func (t *TaskDb) ListTask() {
	var tasks []Tasks.Task
	result := t.db.Find(&tasks)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	for _, task := range tasks {
		fmt.Println(task.ID)
		fmt.Println(task.Title)
		fmt.Println(task.Completed)
	}
}
