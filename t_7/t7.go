package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	myMap := make(map[int]string)

	numGoroutines := 10

	var mutex sync.RWMutex
	var wg sync.WaitGroup

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			key := id
			value := "data is " + strconv.Itoa(id)

			mutex.Lock()
			myMap[key] = value
			mutex.Unlock()
		}(i)
	}

	wg.Wait()

	fmt.Println("map:")
	for key, data := range myMap {
		fmt.Printf("key = %d, value = %s \n", key, data)
	}
}
