package main

import (
	"fmt"
	"time"
)

func main() {
	go goRout()
	time.Sleep(1 * time.Second)
	fmt.Println("Горутина завершилась")
}

func goRout() {
	fmt.Println("Горутина работает")
	panic("panic в горутине")
	fmt.Println("panic не сработала")
}
