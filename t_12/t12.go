package main

import "fmt"

func main() {
	s := []string{"cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]bool)

	for _, item := range s {
		set[item] = true
	}

	uniqueStrings := make([]string, 0, len(set))
	for key := range set {
		uniqueStrings = append(uniqueStrings, key)
	}

	fmt.Println(uniqueStrings)
}
