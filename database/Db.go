package database

import (
	"TodoApp/tasks"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type TaskDb struct {
	db *gorm.DB
}

func NewTaskDb() *TaskDb {
	return &TaskDb{}
}

func (t *TaskDb) LoadDb() error {
	var err error
	t.db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return err
	}
	if err := t.db.AutoMigrate(&tasks.Task{}); err != nil {
		log.Fatal(err)
		return err
	}
	return nil

}

func (t *TaskDb) AddTask(task *tasks.Task) error {
	result := t.db.Create(task)

	if result.Error != nil {
		log.Fatal(result.Error)
		return result.Error
	}
	fmt.Println("Added task successfully")
	return nil
}

func (t *TaskDb) UpdateTask(taskId uint) error {
	result := t.db.Model(&tasks.Task{}).Where("id = ?", taskId).Updates(tasks.Task{Completed: true})
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println("Updated task successfully")
	return nil
}

func (t *TaskDb) DeleteTask(taskId uint) error {
	result := t.db.Model(&tasks.Task{}).Where("id = ?", taskId).Delete(&tasks.Task{})
	if result.Error != nil {
		log.Fatal(result.Error)
		return result.Error
	}
	fmt.Println("Deleted task successfully")
	return nil
}

func (t *TaskDb) ListTask() error {
	var taskslist []tasks.Task
	result := t.db.Find(&taskslist)
	if result.Error != nil {
		log.Fatal(result.Error)
		return result.Error
	}
	for _, task := range taskslist {
		fmt.Println(task.ID)
		fmt.Println(task.Title)
		fmt.Println(task.Completed)
	}
	return nil
}
