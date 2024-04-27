package usecases

import (
	"github.com/ibilalkayy/small-projects/task-clean/entities"
	"github.com/ibilalkayy/small-projects/task-clean/repositories"
)

type AddTaskUseCase struct {
	TaskRepository repositories.TaskRepository
}

func (uc *AddTaskUseCase) Execute(title string) {
	task := entities.Task{ID: 1, Title: title, IsComplete: false}
	uc.TaskRepository.AddTask(task)
}
