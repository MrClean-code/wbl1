package main

import (
	"errors"
	"fmt"
)

var justString string

type err error

func main() {
	someFunc()
}

func createHugeString(size int) (err, string) {
	maxAllowedSize := 1024

	if size <= maxAllowedSize {
		return nil, "1234"
	}

	return errors.New("размер превышает норму"), ""

}

func someFunc() {
	error1, v := createHugeString(1 << 10)

	if error1 != nil {
		fmt.Println(error1)
	}

	if len(v) > 100 {
		justString = v[:100]
		fmt.Println(justString)
	} else {
		justString = v
		fmt.Printf("%s <- Меньше чем нужно символов", v)
	}
}
