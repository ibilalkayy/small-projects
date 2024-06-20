package fetcher

import "fmt"

func fetching(data []string, count chan int) {
	var result int
	for range data {
		result++
	}
	count <- result
}

func Data() {
	data := []string{
		"https://youtube.com",
		"https://linkedin.com",
		"https://twitter.com",
		"https://tiktok.com",
	}

	count := make(chan int)

	go fetching(data, count)

	result := <-count

	fmt.Println(result)
}
