package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{9, 7, 5, 11, 12, 2, 14, 3, 10, 6}

	sort.Ints(arr)

	fmt.Println(arr)
}
