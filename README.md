# Generational Go Runtime Scheduler

## Overview
This project modifies the Go runtime scheduler to divide the `allp` array (which holds processor structures, `P`) into two halves:
- The **first half** is marked as handling **short-running tasks**.
- The **second half** is marked as handling **long-running tasks**.

This enhancement aims to improve scheduling efficiency by allowing better workload distribution and preventing deprioritization of long-running tasks.


![unnamed](https://github.com/user-attachments/assets/7a117629-d278-4329-a01f-2153bc5f92ff)


![unnamed (1)](https://github.com/user-attachments/assets/5f7bb2c8-bbbd-45f8-a080-81848d9bc6fc)


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

![unnamed (2)](https://github.com/user-attachments/assets/1bd10ce8-058b-48ad-a05b-98898c1966bb)

![unnamed (3)](https://github.com/user-attachments/assets/37487243-6480-4f80-b424-604613074c26)
