# Generational Go Runtime Scheduler

## Overview
This project modifies the Go runtime scheduler to divide the `allp` array (which holds processor structures, `P`) into two halves:
- The **first half** is marked as handling **short-running tasks**.
- The **second half** is marked as handling **long-running tasks**.

This enhancement aims to improve scheduling efficiency by allowing better workload distribution and preventing long-running tasks from blocking short-running ones.

## Key Modifications
### Changes in `procresize()`
- `allp` is divided into two groups.
- New fields `shortRunning` and `longRunning` are added to `P`.
- The first `nprocs/2` Ps are flagged as `shortRunning = 1`.
- The last `nprocs/2` Ps are flagged as `longRunning = 1`.

### Potential Next Steps
- Modify scheduling logic to prioritize short-running tasks.
- Implement different scheduling policies for each category.
- Optimize load balancing and task migration between `P`s.

## Setup Instructions
### Clone the Go Source Code
```sh
git clone https://go.googlesource.com/go
cd go/src
```
### Build the modified Go Runtime
```sh
GOROOT=$(pwd) ./bin/go run your_test_program.go
```
