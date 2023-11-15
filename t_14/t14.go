package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a1 interface{} = 1
	var a2 interface{} = "b"
	var a3 interface{} = false
	var a4 interface{} = make(chan float64)

	v1 := reflect.TypeOf(a1)
	v2 := reflect.TypeOf(a2)
	v3 := reflect.TypeOf(a3)
	v4 := reflect.TypeOf(a4)

	fmt.Printf("%s, %s, %s, %s", v1, v3, v2, v4)
}
