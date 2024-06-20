package server

import (
	"fmt"

	"github.com/ibilalkayy/small-projects/golang/rpc/network"
)

func FromNetwork() {
	fmt.Println("Server: The server has received the order")
}

func PasstoProcedure() {
	fmt.Println("Server: Order is passed to procedure")
}

func HandleOrder(task string) {
	FromNetwork()
	PasstoProcedure()
	network.HandleTask(task)
}
