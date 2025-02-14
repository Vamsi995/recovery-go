package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

const switchCount = 1000000 // 1M

func main() {
	runtime.GOMAXPROCS(1) // 1 cpu limit

	ping := make(chan struct{})
	pong := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(2)

	start := time.Now()

	go func() {
		defer wg.Done()
		for i := 0; i < switchCount; i++ {
			<-ping 
			pong <- struct{}{} 
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < switchCount; i++ {
			ping <- struct{}{} 
			<-pong            
		}
	}()

	wg.Wait()
	elapsed := time.Since(start)

	fmt.Printf("Total time for %d context switches: %v\n", switchCount, elapsed)
	fmt.Printf("Average time per switch: %v ns\n", elapsed.Nanoseconds()/switchCount)
}