package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)

const (
	maxRecursion          = 50
	initialGoroutineCount = 100000
	stepSize              = 100000
	maxGoroutineCount     = 500000
)

// recursive function to force stack allocation.
func useStack(depth int) int {
	if depth >= maxRecursion {
		return depth
	}
	return useStack(depth + 1)
}

func main() {
	go func() {
		fmt.Println("pprof server running on http://localhost:6060/debug/pprof/")
		http.ListenAndServe("localhost:8080", nil)
	}()

	for goroutineCount := initialGoroutineCount; goroutineCount <= maxGoroutineCount; goroutineCount += stepSize {
		fmt.Printf("\nRunning with %d Goroutines:\n", goroutineCount)

		// Force an initial GC and print baseline memory statistics.
		runtime.GC()
		var mStart runtime.MemStats
		runtime.ReadMemStats(&mStart)
		fmt.Printf("Before Goroutines:\n")
		fmt.Printf("Alloc = %v KB, Sys = %v KB, NumGC = %v\n",
			mStart.Alloc/1024, mStart.Sys/1024, mStart.NumGC)

		done := make(chan struct{}, goroutineCount)

		for i := 0; i < goroutineCount; i++ {
			go func(id int) {
				useStack(0)
				done <- struct{}{}
			}(i)
		}

		// Wait for all goroutines to finish.
		for i := 0; i < goroutineCount; i++ {
			<-done
		}

		// Force another GC and read memory statistics.
		runtime.GC()
		var mEnd runtime.MemStats
		runtime.ReadMemStats(&mEnd)
		fmt.Printf("\nAfter Goroutines:\n")
		fmt.Printf("Alloc = %v KB, Sys = %v KB, NumGC = %v\n",
			mEnd.Alloc/1024, mEnd.Sys/1024, mEnd.NumGC)

		fmt.Printf("Profiling complete for %d goroutines.\n", goroutineCount)
		time.Sleep(10 * time.Second)
	}

	fmt.Println("Profiling complete. Visit http://localhost:8080/debug/pprof/ for details.")
	time.Sleep(30 * time.Second)

}
