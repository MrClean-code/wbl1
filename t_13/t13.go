package main

import "fmt"

func main() {
	x := 5
	y := 7

	x, y = y, x

	fmt.Printf("x = %d, y = %d\n", x, y)
}
