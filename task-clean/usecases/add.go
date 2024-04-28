package usecases

import "github.com/ibilalkayy/small-projects/task-clean/entities"

func AddTask(title, description string) *entities.Task {
	adding := entities.Task{
		Title:       title,
		Description: description,
		Completed:   false,
	}
	return &adding
}
