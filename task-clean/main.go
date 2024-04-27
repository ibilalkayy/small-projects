package main

import (
	"fmt"

	"github.com/ibilalkayy/small-projects/task-clean/repositories"
	"github.com/ibilalkayy/small-projects/task-clean/usecases"
)

func main() {
	taskRepository := repositories.InMemoryTaskRepository{}
	addTaskUseCase := usecases.AddTaskUseCase{TaskRepository: &taskRepository}
	getTaskUseCase := usecases.GetTaskUseCase{TakeRepository: &taskRepository}

	addTaskUseCase.Execute("This is the first title")
	addTaskUseCase.Execute("This is the second title")

	tasks := getTaskUseCase.Execute()
	for _, task := range tasks {
		fmt.Printf("ID: %d, Title: %s, Is completed: %t\n", task.ID, task.Title, task.IsComplete)
	}
}
