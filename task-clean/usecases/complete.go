package usecases

import "github.com/ibilalkayy/small-projects/task-clean/entities"

func CompleteTask(tasks *entities.Task) {
	tasks.Completed = true
}
