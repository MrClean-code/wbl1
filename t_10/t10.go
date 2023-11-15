package main

import (
	"fmt"
	"sort"
)

func main() {
	masTemp := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	myMap := make(map[int][]float64)

	sort.Float64s(masTemp)

	minEl := int(masTemp[0])
	maxEl := int(masTemp[len(masTemp)-1])

	minDiap := (minEl / 10) * 10
	maxDiap := (maxEl / 10) * 10

	for temp := minDiap; temp <= maxDiap; temp += 10 {
		myMap[temp] = []float64{}
	}

	for _, temp := range masTemp {
		groupKey := (int(temp) / 10) * 10
		myMap[groupKey] = append(myMap[groupKey], temp)
	}

	for key, values := range myMap {
		fmt.Printf("%d: %v \n", key, values)
	}
}
