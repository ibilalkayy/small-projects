package server

import (
	"fmt"

	"github.com/ibilalkayy/small-projects/todo-clean/rpc/network"
)

func OrderReceived() bool {
	fmt.Println("Server: Order is received")
	return true
}

func GiveTaskToProcedure(task string) bool {
	network.OrderIsTransferring()
	received := OrderReceived()
	if received {
		fmt.Println("Server: The task is given to the procedure")
		network.OrderCompleted(task)
		return true
	} else {
		fmt.Println("Server: The task is not received to the procedure")
		return false
	}
}

func HandleOrder(task string) {
	success := GiveTaskToProcedure(task)
	if !success {
		fmt.Println("Server: Failed to handle the order")
	}
}
