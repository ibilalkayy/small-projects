package network

import (
	"fmt"

	"github.com/ibilalkayy/small-projects/todo-clean/rpc/procedure"
)

func FromClient() {
	fmt.Println("Network: Order is received from the client")
}

func PassToServer() {
	fmt.Println("Network: The order is passed to the server")
}

func HandleTask(task string) {
	if task == "Rice" {
		FromClient()
		PassToServer()
		procedure.HandleProcedure()
	} else {
		fmt.Println("Network: Unknown task")
	}
}
