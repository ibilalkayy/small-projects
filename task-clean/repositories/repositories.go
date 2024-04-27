package repositories

import "github.com/ibilalkayy/small-projects/task-clean/entities"

type TaskRepository interface {
	AddTask(entities.Task)
	GetTask() []entities.Task
}

type InMemoryTaskRepository struct {
	Tasks []entities.Task
}

func (repo *InMemoryTaskRepository) AddTask(task entities.Task) {
	repo.Tasks = append(repo.Tasks, task)
}

func (repo *InMemoryTaskRepository) GetTask() []entities.Task {
	return repo.Tasks
}
