package add_square

import "fmt"

func calculate(num int, squareChan chan int) {
	result := num + num
	squareChan <- result
}

func TakeSquare() {
	squareChan := make(chan int)

	numbers := []int{1, 2, 3, 4, 5}

	for _, num := range numbers {
		go calculate(num, squareChan)
	}

	for range numbers {
		square := <-squareChan
		fmt.Println("Square: ", square)
	}
}
