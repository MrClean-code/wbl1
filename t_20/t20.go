package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun"

	sl := strings.Split(str, " ")

	for i := len(sl) - 1; i >= 0; i-- {
		fmt.Print(sl[i], " ")
	}

}
