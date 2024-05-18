package client

import (
	"fmt"

	"github.com/ibilalkayy/small-projects/todo-clean/rpc/server"
)

func OrderFood() {
	fmt.Println("Client: Ordering Rice")
	server.HandleOrder("Rice")
}
