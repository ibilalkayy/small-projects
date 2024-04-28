package main

import (
	"fmt"
	"log"

	"github.com/ibilalkayy/small-projects/calculate-clean/entities"
	"github.com/ibilalkayy/small-projects/calculate-clean/usecases"
)

func main() {
	operand1 := 6
	operand2 := 6

	op := &entities.Operation{Operand1: operand1, Operand2: operand2}

	addResult := usecases.Add(op)
	subtractResult := usecases.Subtract(op)
	multiplyResult := usecases.Multiply(op)
	divideResult, err := usecases.Divide(op)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(addResult)
	fmt.Println(subtractResult)
	fmt.Println(multiplyResult)
	fmt.Println(divideResult)
}
