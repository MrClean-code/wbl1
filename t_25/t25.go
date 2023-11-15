package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func sleep(seconds int) {
	defer wg.Done()

	done := make(chan bool)
	go func() {
		select {
		case <-time.After(time.Duration(seconds) * time.Second):
			done <- true
		}
		close(done)
	}()

	<-done
}

func main() {
	fmt.Println("Начало работы")
	wg.Add(1)
	go sleep(5)
	wg.Wait()
	fmt.Println("Прошло 5 секунд")
}
