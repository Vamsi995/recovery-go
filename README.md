# Towards a Self Healing Runtime for Go

## Vision Statement

Concurrency is a core feature of the Go programming language, enabling developers to efficiently utilize multi-core processors. However, common concurrency issues such as deadlocks, race conditions, and goroutine leaks can severely impact system reliability. The goal of this project is to first develop a deep understanding of Go’s concurrency model, benchmark different concurrency mechanisms, and ultimately build a self-healing runtime tool capable of detecting and recovering from concurrency issues dynamically.

## Goals & Objectives

- Gain proficiency in Go’s concurrency model, including goroutines, channels, and synchronization primitives.
- Benchmark the performance of different concurrency mechanisms (e.g., mutexes, channels, atomic operations).
- Develop an automated system to detect and recover from concurrency issues like deadlocks, race conditions, and goroutine leaks.
- Evaluate the effectiveness of the self-healing runtime through empirical analysis.

## Project Plan

### ✅ | Phase 1: Learn Go Programming Language

### ⬜️ | Phase 2: Benchmarking Go’s Concurrency Mechanisms
- Contrast Go Concurrency with other languages like C++, Java
  - Thread/Goroutine Creation Time:	The language runtime manages thread spawning and scheduling.
  - Context Switching Overhead:	Go's cooperative scheduler vs. OS-managed preemptive scheduling (C++/Java).
  - Synchronization Performance: How runtime-managed mutexes, atomic operations, and channels perform.
  - Memory Usage (Heap & Stack Growth):	Goroutine stack management (grows dynamically) vs. fixed stacks in C++/Java.
  - Garbage Collection Impact: Go & Java have GC
  - Scalability Test (100K Goroutines vs Threads): M:N Goroutine scheduling vs OS threading scalability.

### ⬜️ | Phase 3: Designing a Self-Healing Runtime
- **Deadlock Detection & Recovery**
  - Use timeouts and forced unlocking strategies to resolve deadlocks
  - Can we develop a runtime ebpf application here that monitors Go's goroutine threads?
- **Goroutine Leak Detection**
  - Implement the self-healing runtime as a Go library.





