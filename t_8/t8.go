package main

import (
	"fmt"
)

func main() {
	var number int64 = 64
	index := 0
	bit := true // 1 - true; 0 - false

	result := setBit(number, index, bit)
	fmt.Printf("Исходное число: %d\n", number)
	fmt.Printf("Результат: %d\n", result)
}

func setBit(n int64, i int, value bool) int64 {
	if value {
		return n | (1 << i)
	} else {
		return n &^ (1 << i)
	}
}
