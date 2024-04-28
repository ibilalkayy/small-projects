package usecases

import (
	"errors"

	"github.com/ibilalkayy/small-projects/calculate-clean/entities"
)

func Divide(op *entities.Operation) (int, error) {
	if op.Operand2 == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return op.Operand1 / op.Operand2, nil
}
