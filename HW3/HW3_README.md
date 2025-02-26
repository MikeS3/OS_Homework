# HW3 Concurrent Linked list Queue Implementation Comparison

Comparing the execution times of A ticket lock and Compare and Swap Spin Lock

## The program overview

- **TicketLock** Structure with ticket and turn variables, and lock and unlock function
- **CAS Spin Lock** Structure with a flag variable and lock and unlock functions

## How It Works

- **TicketLock**
  - Lock, uses the fetch and add instruction equivalent to decrement the ticket counter
  - Spins after setting the lock
  - Unlock imcrements the counter with the Go atomic add primitive

- **CASSpinLock**
  - Has a flag variable that is set to 0 or 1 to show the lock availability
  - Uses the Compare and Swap atomic primitive to see the lock availability and grab the lock if available
  - Unlocks by setting the flag to 0

- **Benchmarking**
  - Create 32 Go routines for the ticket lock and time how long it takes to execute
  - Create 32 Go routines after the ticket locks have run and been timed, then time how long it takes for the CAS Spin lock Go routines to execute

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) is installed so you can run program.

### Cloning the Repository

Clone the repository using the following commands:

```bash
git clone https://github.com/MikeS3/OS_Homework.git
cd OS_Homework
cd HW2
```
### Running the Code

To run the code enter this command
```go run HW2.go```

### Results 
The expected output should look something like this, with times dependent on the computer running. We found the ticket lock to take longer than the
compare and swap spin lock

Expected output:
Ticket Lock (with 32 goroutines) took: 617.536692ms
CAS Spin Lock (with 32 goroutines) took: 518.150893ms
