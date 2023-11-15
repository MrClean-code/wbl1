package main

import (
	"fmt"
	"time"
)

// 2 вид решения
func main() {
	fmt.Print("Введите время: ")
	var timeC int
	fmt.Scan(&timeC)

	timer := time.NewTimer(time.Duration(timeC) * time.Second)

	channel := make(chan int)

	number := 5349
	go func() {
		defer close(channel)

		for {
			channel <- number
			fmt.Printf("Запись %d \n", number)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for res := range channel {
			fmt.Printf("Чтение %d \n", res)
			fmt.Println()
		}
	}()

	select {
	case <-timer.C:
		fmt.Println("Программа завершена.")
	}
}
