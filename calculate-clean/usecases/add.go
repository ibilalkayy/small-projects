package usecases

import "github.com/ibilalkayy/small-projects/calculate-clean/entities"

func Add(op *entities.Operation) int {
	return op.Operand1 + op.Operand2
}
