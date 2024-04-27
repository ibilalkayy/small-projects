package usecases

import (
	"github.com/ibilalkayy/small-projects/task-clean/entities"
	"github.com/ibilalkayy/small-projects/task-clean/repositories"
)

type GetTaskUseCase struct {
	TakeRepository repositories.TaskRepository
}

func (uc *GetTaskUseCase) Execute() []entities.Task {
	return uc.TakeRepository.GetTask()
}
