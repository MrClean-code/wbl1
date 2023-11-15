package main

import (
	"fmt"
)

func main() {
	word := "да уж ну и день"
	revStr := reverseStr(word)
	fmt.Println(revStr)
}

func reverseStr(word string) interface{} {

	sl := []rune{}
	for _, item := range word {
		sl = append(sl, item)
	}

	for i, j := 0, len(sl)-1; j >= i; i, j = i+1, j-1 {
		sl[i], sl[j] = sl[j], sl[i]
	}

	return string(sl)
}
