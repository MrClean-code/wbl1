package main

import (
	"fmt"
	"sync"
)

var exitCh = make(chan string)
var wg sync.WaitGroup

func main() {
	fmt.Println("start")
	wg.Add(1)
	go goRoutine()
	close(exitCh)
	wg.Wait()
	fmt.Println("end")
}

func goRoutine() {
	defer wg.Done()

	for {
		select {
		case <-exitCh:
			fmt.Println("stop goroutine ")
			return
		}
	}
}
