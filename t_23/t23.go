package main

import (
	"fmt"
	"log"
)

func main() {
	sl := []float64{32.0, 33.0}

	fmt.Print("i element: ")
	var i int
	fmt.Scan(&i)

	sl = deleteEl(i, sl)
	if sl == nil {
		log.Panicf("нету значения по такому индексу %d", i)
	}
	fmt.Println(sl)
}

func deleteEl(i int, sl []float64) []float64 {
	if i <= len(sl)-1 {
		sl = append(sl[:i], sl[i+1:]...)
		return sl
	}

	return nil
}
