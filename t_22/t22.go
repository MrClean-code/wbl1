package main

import (
	"fmt"
	"math"
)

func main() {
	a := math.Pow(2, 20)
	b := math.Pow(2, 20)

	rD := div(a, b)
	rP := proizv(a, b)
	rR := razn(a, b)
	rS := sum(a, b)

	fmt.Println(rD)
	fmt.Println(rP)
	fmt.Println(rR)
	fmt.Println(rS)
}

func sum(a, b float64) float64 {
	return a + b
}

func razn(a, b float64) float64 {
	return a - b
}

func proizv(a, b float64) float64 {
	return a * b
}

func div(a, b float64) float64 {
	return a / b
}
