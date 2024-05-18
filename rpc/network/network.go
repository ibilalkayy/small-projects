package network

import (
	"fmt"

	"github.com/ibilalkayy/small-projects/todo-clean/rpc/procedure"
)

func OrderIsTransferring() {
	fmt.Println("Network: Order is transferred")
}

func OrderCompleted(task string) {
	procedure.HandleTask(task)
}
