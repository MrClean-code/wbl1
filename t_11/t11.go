package main

import "fmt"

func main() {
	set1 := map[string]bool{"apple": true, "banana": true, "cherry": true}
	set2 := map[string]bool{"banana": true, "date": true, "fig": true}

	result := createSet(set1, set2)

	fmt.Println(result)
}

func createSet(set1, set2 map[string]bool) map[string]bool {
	result := make(map[string]bool)

	for item := range set1 {
		if set2[item] {
			result[item] = true
		}
	}

	return result
}
