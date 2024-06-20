package add_square

import "fmt"

func square(num int, squareChan chan int) {
	result := num * num
	squareChan <- result
}

func add(num int, addChan chan int) {
	result := num + num
	addChan <- result
}

func Combined() {
	addChan := make(chan int)
	squareChain := make(chan int)

	go add(4, addChan)
	go square(4, squareChain)

	addResult := <-addChan
	squareResult := <-squareChain

	fmt.Println(addResult)
	fmt.Println(squareResult)
}
