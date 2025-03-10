package main

// import ("fmt"
// "reflect")

// func main() {
//     ch1 := make(chan int)
//     ch2 := make(chan int)

//     go func() {
//         <-ch1
//         ch2 <- 1
//     }()

//     <-ch2
//     ch1 <- 1
//     //<-ch2 // Deadlock: Circular Wait
// }

// func main() {
//     ch1 := make(chan int)
//     ch1 <- 1
// }

// func main() {
//     ch := make(chan int)
//     // fmt.Println("Type of x:", reflect.TypeOf(ch).Kind())

//     // if reflect.TypeOf(ch).Kind().String() == "chan" {
//     //     fmt.Println("yes")
//     // }

//     // reflect.TypeOf(ch)
//     // go func() { ch <- 1 }()
//     <-ch
// }

// package main

import (
	"os"
	"runtime/trace"
	// "fmt"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	// "time"
)

// MatrixSize defines the size of the matrices to multiply
const matrixSize = 500 // Adjust for higher CPU load

// numWorkers defines the number of parallel goroutines
const numWorkers = 8 // Adjust based on GOMAXPROCS

// generateMatrix creates a random NxN matrix
func generateMatrix(n int) [][]float64 {
	matrix := make([][]float64, n)
	for i := range matrix {
		matrix[i] = make([]float64, n)
		for j := range matrix[i] {
			matrix[i][j] = rand.Float64() * 10
		}
	}
	return matrix
}

// multiplyMatrices performs matrix multiplication
func multiplyMatrices(A, B [][]float64) [][]float64 {
	N := len(A)
	C := make([][]float64, N)
	for i := range C {
		C[i] = make([]float64, N)
		for j := range C[i] {
			sum := 0.0
			for k := 0; k < N; k++ {
				sum += A[i][k] * B[k][j]
			}
			C[i][j] = sum
		}
	}
	return C
}

// func worker(id int, wg *sync.WaitGroup) {
// 	defer wg.Done()
	
// 	fmt.Printf("Worker %d completed in %v\n", id, elapsed)
// }


func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	// fibonacci(40) // Simulated CPU-bound work
	// fmt.Println(out)
	// A := generateMatrix(matrixSize)
	// B := generateMatrix(matrixSize)
	// // start := time.Now()
	// _ = multiplyMatrices(A, B) // Compute but discard result
	// elapsed := time.Since(start)
	result := 0
	for i := 0; i < 1e9; i++ { result += i }

}

func worker_short(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	// fibonacci(10) // Simulated CPU-bound work
	// fmt.Println(out)
	// A := generateMatrix(matrixSize)
	// B := generateMatrix(matrixSize)
	// // start := time.Now()
	// _ = multiplyMatrices(A, B) // Compute but discard result
	// elapsed := time.Since(start)
	result := 0
	for i := 0; i < 1e6; i++ { result += i }

}

func main() {
    runtime.GOMAXPROCS(4)
	f, err := os.Create("trace_preempt.out")
	if err != nil {
		fmt.Println("Error creating trace file:", err)
		return
	}
	defer f.Close()

	trace.Start(f)
	defer trace.Stop()

	// start := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go worker(i, &wg)
		for j := 0; j < 100; j++ {
			wg.Add(1)
			go worker_short(j, &wg)
		}
	}
	for j := 0; j < 20000; j++ {
		wg.Add(1)
		go worker_short(j, &wg)
	}
	wg.Wait()

	// fmt.Println(time.Since(start))
}
