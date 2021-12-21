package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	data := 0
	wg.Add(1)
	go func() {
		defer wg.Done()
		data++
	}()

	wg.Wait()
	fmt.Printf("Number: %v\n", data)
}
