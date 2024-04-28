package main

import (
	"fmt"
	"os"

	"github.com/ibilalkayy/small-projects/task-clean/entities"
	"github.com/ibilalkayy/small-projects/task-clean/usecases"
)

func main() {
	tasks := []*entities.Task{}

	for {
		fmt.Println("Task Manager")
		fmt.Println("1. Add Task")
		fmt.Println("2. View Task")
		fmt.Println("3. Complete Task")
		fmt.Println("4. Exit")

		var choice int
		fmt.Printf("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var title, description string
			fmt.Print("Enter your title: ")
			fmt.Scanln(&title)
			fmt.Print("Enter your description: ")
			fmt.Scanln(&description)
			tasks = append(tasks, usecases.AddTask(title, description))
		case 2:
			fmt.Println("Tasks:")
			for i, task := range tasks {
				fmt.Printf("%d: Title: %s, Description: %s, Completed: %t\n", i+1, task.Title, task.Description, task.Completed)
			}
		case 3:
			var index int
			fmt.Print("Enter the index of the task to mark as completed: ")
			if index >= 1 && index <= len(tasks) {
				usecases.CompleteTask(tasks[index-1])
			} else {
				fmt.Println("invalid index")
			}
		case 4:
			fmt.Println("Exiting the task manager")
			os.Exit(0)
		default:
			fmt.Println("invalid choice")
		}
	}
}
