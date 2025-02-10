# HW1 Python Threads vs Go GoRoutines

Python thread, implementation of IPC is tested against GO's GoRoutines 

## The program overview

- **Producer:** Writes numbers (1 to 5) into the pipe.
- **Consumer:** Reads numbers from the pipe and prints them until the pipe is closed.

## How It Works

- **Producer:**
  - Iterates through numbers 1 to 5.
  - Writes each number to the pipe.
  - We have a 500ms delay between writes to see processes better in both python and Go.
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
- [Python](https://www.python.org/downloads/) is installed so you can run the python program
- [bc] Basic calculator is intalled in your terminal so the bash script can calculate the time taken by each program

### Cloning the Repository

Clone the repository using the following commands:

```bash
git clone https://github.com/MikeS3/OS_Homework.git
cd OS_Homework
cd HW1
```
### Running the Code

To run the code enter this command
First ensure the bash script has execute permissions by running ```chmod +x benchmark.sh```

next run the benchmark program which will run and time both the python threadi mplementation and the Go GoRoutine Implementation
by running

```
./benchmark.sh
```

### Results 
Both Programs are run and the the time it takes for them to execute is printed

Expected output:
Starting Implementation with Go routines...
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
Starting Python Threads Implementation ...
Consumer: 1
Producer: 1
Producer: 2
Consumer: 2
Producer: 3
Consumer: 3
Producer: 4
Consumer: 4
Producer: 5
Consumer: 5
Goroutine Implementation: 2.696169834 seconds
Python Thread Implementation: 2.534811611 seconds
