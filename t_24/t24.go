package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) *Point {
	return &Point{x: x, y: y}
}
func (p *Point) diffPoint() float64 {
	x := math.Abs(p.x)
	y := math.Abs(p.y)
	diff := 0.0
	if x > y {
		diff = x - y
	} else {
		diff = y - x
	}
	return diff
}

func main() {
	point := NewPoint(12.3, 30.4)
	diff := point.diffPoint()
	fmt.Println(diff)
}
