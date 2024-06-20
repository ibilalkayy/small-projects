package client

import (
	"fmt"

	"github.com/ibilalkayy/small-projects/golang/rpc/network"
)

func PlacedAnOrder() {
	fmt.Println("Client: Rice is ordered")
	network.HandleTask("Rice")
}
