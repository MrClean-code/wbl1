package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg2 sync.WaitGroup

func main() {
	wg2.Add(1)
	go goRoutin()
	wg2.Wait()
	fmt.Println("Горутина завершилась")
}

func goRoutin() {
	defer wg2.Done()
	fmt.Println("Горутина работает")
	runtime.Goexit()
	fmt.Println("остановка горутины не произошло")
}
