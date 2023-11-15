package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	channel1 := make(chan int, 5)
	channel2 := make(chan int, 5)
	numbers := []int{2, 4, 6, 8, 10}

	for _, number := range numbers {
		wg.Add(1)
		go addChannelOne(number, &wg, channel1)
		wg.Add(1)
		go addChannelSecond(number, &wg, channel2)
	}

	wg.Wait()
	close(channel1)
	close(channel2)

}

func addChannelSecond(number int, wg *sync.WaitGroup, channel2 chan int) {
	defer wg.Done()

	channel2 <- number * 2

	fmt.Printf("Запись в 2 канал = %d \n", <-channel2)
}

func addChannelOne(number int, wg *sync.WaitGroup, channel1 chan int) {
	defer wg.Done()

	channel1 <- number

	fmt.Printf("Запись в 1 канал = %d \n", <-channel1)
}
