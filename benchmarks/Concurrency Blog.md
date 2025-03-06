# Concurrency in Go Runtime

Go is a programming language that utilizes lightweight concurrency units called goroutines to achieve concurrency. Unlike traditional programming languages such as C++ and Java, which rely on a threading-based concurrency model, Go follows the Communicating Sequential Processes (CSP) concurrency model. The Go runtime is responsible for scheduling the execution of these goroutines efficiently.


## CSP (Communicating Sequential Processes) Concurrency Model

Concurrency introduces challenges in managing shared memory across different processes or threads. The non-deterministic execution order of concurrent tasks makes debugging and reproducing errors difficult. Traditional thread-based concurrency models address these issues using mutexes and locks. However, these mechanisms can lead to data races, where multiple threads access a shared resource simultaneously, or deadlocks, where threads wait indefinitely for resource access.

The CSP concurrency model mitigates these problems by eliminating direct memory sharing. Instead of threads sharing memory and requiring synchronization mechanisms, CSP facilitates communication between concurrent processes through channels. Channels function similarly to Linux pipes, enabling safe data exchange without explicit locking.

A core philosophy of Go's concurrency model is:

> "Don't communicate by sharing memory; instead, share memory by communicating."

This paradigm simplifies concurrent programming by reducing synchronization complexity and minimizing race conditions, making Go an excellent choice for building scalable and efficient concurrent applications.




## Concurrency vs Parallelism - What is the difference?

Concurrency is managing many things at once, whereas parallelism is doing many things at once. The key difference being with a single process I can execute two processes concurrently but I cannot execute them parallely because I will be needing another processing unit to do both the tasks at the same time.


## Diving Deeper into Go's Concurrency

Go's concurrency is handled by the Go Runtime. Most of the concurrency in Go is through the paradigm of go routines that implement CSP model. The way the runtime handles the go routines is quite interesting and we will dive deeper into the internal workings of the runtime scheduler of these go routines.


### Go Runtime Scheduler for Go Routines

#### Go's Scheduler Model

Go uses go routines and channels as concurrency primitives. Let's focus on go routines. Go routines are analagous to coroutines in concurrency programming model where a piece of execution unit can be run, paused and resumed. Unlike the OS threads which does not give us control on the scheduling/preempting process of the threads, the go routines are implemented in the user space. The go runtime is responsible for scheduling these go routines and managing their execution during the programs lifetime. The Go runtime sits between the operation system and our application code, and orchestrates the go routines for us abstracting away the operating system kernel.

![alt text](/images/Runtime_Hierarchy.png)

Go's runtime allows us to spawn as many go routines as we want, despite the hardware limitation. The problem that the runtime solves is how to manage these **n** number of go routines **G** on a machine with **m** hardware processors.

![alt text](/images/Go%20Routine%20Mapping.png)



Each hardware processor gets a single thread assigned at any point of time, so at any given instant in our program we will have **m** number of OS threads executing on **m** hardware processors. Note that we can have multiple kernel threads running, but only **m** number of threads will be in the running state at any given point.

![alt text](/images/Mapping.png)

Go runtime is responsible for running go routines on operating system threads i.e on our hardware. This is called user-space scheduling. There are multiple ways of user-space scheduling out there, but go follows something called **n:m** scheduling. This allows for the program to have more number of go routines than the operating system threads. These go routines are multiplexed onto these OS threads. 

![alt text](/images/NMScheduling.png)
    
#### Keeping Track of Go Routines to Run - Global Run Queues

Global Run Queues are used to keep track of the Go Routines. After creating a go routine, the runtime scheduler appends it to the global runqueue. From here the assignment is bidirectional, i.e the runtime can assign different threads, or the threads can steal go routines from the global runqueue.


![alt text](/images/GlobalRunQueue.png)

Now to for the threads to access the global runqueue, we need a lock on the queue to synchorize access. But the problem with the lock is that if we have a go routine that runs for a less amount of time, then after finishing  it the thread has to acquire the global runqueue lock and steal more go routines from the queue. This way there would be a contention overhead on the access to the global runqueue. 

![alt text](/images/lock.png)


To avoid this contention overhead, a concept of distributed local runqueues were introduced. This way the thread need not keep a lock to access the next go routines to run after it finished a go routine.

![alt text](/images/LocalQueue.png)














GOMAXPROCS limits the number of operating system threads running on the kernel space. Note that there is no limit on the number of OS threads that are blocked because of a go routine, but those do not count against the GOMAXPROCS limit.


Go scheduler implements work stealing threads, i.e if a thread is idle, it can steal from other threads which have a lot of go routines to process. But this way it has to search through an unbounded number of threads. To avoid this notion, go runtime has another overlay over threads called processors. Processors are heap allocated data-structure, and they hold whatever the state the thread was holding before like the local runqueue. Number of processors = GOMAXPROCS. This way we only have to check all the processors for work stealing insteaad of an unbounded number of threads. 

![alt text](/images/processor.png)



![alt text](/images/RuntimeProcessor.png)

#### Fairness & Preemption

FIFO local runqueues expose a problem of fairness, which deals with a single go routine taking up too much time on the thread, i.e a resource hog which can starve all the other go routines in the local runqueue. The way to solve this is through preemption just like OS thread preemption. There are two types Co-operative Preemption, & Non Co-operative Preemption. In co-operative preemption we trust the execution unit to give up control after some period of time. But this is not always the case, so non-cooperative preemption is used in such scenarios where we use the concepts of timeslice, and preempt the go routine after running for this timeslice. In the Go Runtime this timeslice is of 10ms, and is initiated by SIGURG signal called by the sysmon daemon go routine periodically.


#### Work Stealing of Go Routines
First check local runqueue, then if empty, check global runqueue, then take your share of go routines, if global runqueue is empty, steal from netpoller (here netpoller is a buffer of go routines waiting on network IO), if netpoller is empty, steal from other processor, pick random P -> if work available, steal half of it. 