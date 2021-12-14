package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	numbers := generateNumbers(100000000, 200)
	t := time.Now()
	sum := add(numbers)
	fmt.Printf("Sequential add: Sum: %v  Time: %v\n", sum, time.Since(t))

	t = time.Now()
	sum = concurrentAdd(numbers)
	fmt.Printf("Concurrent add: Sum: %v Time: %v\n", sum, time.Since(t))
}

func generateNumbers(n int, threshold int) []int {
	v := make([]int, n)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		v[i] = rand.Intn(threshold)
	}

	return v
}

func add(numbers []int) int64 {
	var res int64
	for _, n := range numbers {
		res += int64(n)
	}

	return res
}

func concurrentAdd(numbers []int) int64 {
	// check all cores on machine
	processors := runtime.NumCPU()
	runtime.GOMAXPROCS(processors)

	var sum int64
	sizeOfParts := len(numbers) / processors

	var wg sync.WaitGroup

	for i := 0; i < processors; i++ {
		// Divide the input into parts
		start := i * sizeOfParts
		end := start + sizeOfParts
		part := numbers[start:end]

		wg.Add(1)
		go func(nums []int) {
			defer wg.Done()
			partSum := int64(0)

			for _, n := range nums {
				partSum += int64(n)
			}
			atomic.AddInt64(&sum, partSum)

		}(part)
	}

	wg.Wait()
	return sum
}
