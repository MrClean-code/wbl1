package main

import (
	"fmt"
	"sort"
)

func main() {

	arr := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	findEl := 6
	result := binarySearch(arr, findEl)

	if result != -1 {
		fmt.Printf("Элемент %d найден", findEl)
	} else {
		fmt.Printf("Элемент %d не найден", findEl)
	}
}

func binarySearch(arr []int, target int) int {
	sort.Ints(arr)

	index := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= target
	})

	if index < len(arr) && arr[index] == target {
		return index
	} else {
		return -1
	}
}
