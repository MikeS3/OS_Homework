# HW0 Process and Inter-Process Communication (IPC)

A simple implementation of inter-process communication (IPC) using pipes in Go. Pipe used is for the concurrency(IPC) to work with the two processes(producer and consumer)

## The HW overview

- **Producer:** Writes numbers (1 to 5) into the pipe.
- **Consumer:** Reads numbers from the pipe and prints them until the pipe is closed.
- **Synchronization:** Uses `sync.WaitGroup` to ensure both goroutines are complete before the program exits.

## How It Works

- **Producer:**
  - Iterates through numbers 1 to 5.
  - Writes each number to the pipe.
  - We have a 500ms delay between writes to see processes better.
  - Closes the pipe after writing all values.
  
- **Consumer:**
  - Reads integers from the pipe.
  - Prints each received number.
  - Stops reading when the pipe is closed.

- **Main Function:**
  - Creates a pipe (`os.Pipe()`).
  - Launches producer and consumer goroutines.
  - Waits for both to complete using `sync.WaitGroup`.

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) is installed so you can run program.

### Cloning the Repository

Clone the repository using the following commands:

```bash
git clone https://github.com/MikeS3/OS_Homework.git
cd OS_Homework
cd HW0
```
### Running the Code

To run the code enter this command

```
go run HW0.go
```
### Results 
Produce successfully gives the 5 numbers and the consumer gets them and prints them

Expected output:

Producer: 1

Consumer: 1

Producer: 2

Consumer: 2

Producer: 3

Consumer: 3

Producer: 4

Consumer: 4

Producer: 5

Consumer: 5


