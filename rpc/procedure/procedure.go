package procedure

import (
	"fmt"
)

func Received() {
	fmt.Println("Procedure: Task is received to be completed")
}

func MakingRice() {
	fmt.Println("Procedure: Making Rice and it is now ready")
}

func HandleTask(task string) {
	if task == "Rice" {
		Received()
		MakingRice()
	} else {
		fmt.Println("Procedure: Unknown task")
	}
}
