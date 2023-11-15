package main

import (
	"fmt"
	"time"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	channel := make(chan int, 5)
	sumCh := make(chan int, 1)

	for _, number := range numbers {
		go square(number, channel)
	}

	go func() {
		sum := 0
		for res := range channel {
			sum += res
		}
		sumCh <- sum
		close(sumCh)
	}()
	time.Sleep(300 * time.Millisecond)

	close(channel)

	fmt.Println(<-sumCh)
}

func square(number int, chanel chan int) {
	sq := number * number
	chanel <- sq
}
