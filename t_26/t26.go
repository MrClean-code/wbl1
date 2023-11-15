package main

import (
	"fmt"
	"strings"
)

func main() {
	inputStrings := []string{
		"abcd",
		"GoLang",
	}

	for _, str := range inputStrings {
		fmt.Printf("%s — %v \n", str, allCharactersUnique(str))
	}
}

func allCharactersUnique(s string) bool {
	charSet := make(map[rune]bool)
	for _, char := range strings.ToLower(s) {
		//fmt.Println(charSet[char])
		if charSet[char] {
			return false
		}
		charSet[char] = true
	}
	return true
}
