package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// benchmarkGoroutines measures the time taken to spawn `n` goroutines
func benchmarkGoroutines(n int) time.Duration {
	var wg sync.WaitGroup
	wg.Add(n)

	start := time.Now()
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
		}()
	}
	wg.Wait()
	elapsed := time.Since(start)

	return elapsed
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // Use all available CPUs

	// Different counts of goroutines to test
	testCases := []int{100, 1000, 10_000, 100_000}

	fmt.Println("Benchmarking goroutine creation time:")
	for _, n := range testCases {
		duration := benchmarkGoroutines(n)
		fmt.Printf("Created %d goroutines in %v\n", n, duration)
	}
}

