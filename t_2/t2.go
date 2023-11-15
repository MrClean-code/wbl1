package main

import (
	"fmt"
	"time"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	chanel := make(chan int, 5)

	for _, number := range numbers {
		go square(number, chanel)
	}
	time.Sleep(300 * time.Millisecond)

	close(chanel)

	for res := range chanel {
		fmt.Println(res)
	}

}

func square(number int, chanel chan int) {
	sq := number * number
	chanel <- sq
}
