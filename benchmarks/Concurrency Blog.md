# Concurrency in Go Runtime

Go is a programming language that utilizes light weight concurrency units called go routines to achieve concurrency. Like other programming lanugages the runtime is responsible for scheduling the execution of these goroutines. Go follows the CSP concurrency model, unlike other traditional programming lanugages like C++, Java which use threading based concurrency model.



## CSP (Communicating Sequential Processes) Concurrency Model

Concurrency has the traditional problem of sharing memory across different processes/threads. Memory sharing is a non-deterministic problem because the execution of the concurrent tasks have no pre-determined order, which infact makes it harder to debug, reproduce errors. The thread based concurrency model solves most of these issues using mutexes/locks. 

Traditional Programming models have concurrency with shared memory and using locks, but in CSP each process/thread does not share memory but instead communicate with each other using channels which are like pipes in Linux.

dont communicate by sharing memory, but share memory by communicating


## Concurrency vs Parallelism - What is the difference?




## Diving Deeper into Go's Concurrency


