package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	var wg sync.WaitGroup

	fmt.Print("Кол-во пользователей: ")
	var workers int
	fmt.Scan(&workers)

	channel := make(chan string)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT)

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go worker(i, channel, &wg)
	}

	data := "0123456789"
	for {
		select {
		case <-quit:
			fmt.Println("Ctrl+C ---> остановка ")
			wg.Wait()
			close(channel)
			return
		case channel <- data:
		}
	}

}

func worker(workerID int, channel <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for data := range channel {
		fmt.Printf("Worker %d read: %s\n", workerID+1, data)
	}

}
