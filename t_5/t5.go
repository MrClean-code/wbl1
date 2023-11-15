package main

import (
	"fmt"
	"time"
)

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
			fmt.Printf("Запись %d прошло за 500мс \n ", number)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	for {
		select {
		case data := <-channel:
			fmt.Printf("Чтение %d \n", data)
		case <-timer.C:
			fmt.Println("Программа завершена.")
			return
		}
	}
}
