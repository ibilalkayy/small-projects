package procedure

import "fmt"

func FromServer() {
	fmt.Println("Procedure: The order is received from the server")
}

func MakingRice() {
	fmt.Println("Procedure: Cooking the rice and now it is ready")
}

func HandleProcedure() {
	FromServer()
	MakingRice()
}
