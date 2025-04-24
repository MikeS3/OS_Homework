# HW4 Concurrent Queues: Michael and Scott vs. Lock-Free Queues  
Comparing the execution times of Michael and Scott's and Lock-Free concurrent queues.

## **Program Overview**  
This program implements and benchmarks two types of queues:
1. **Michael and Scott's Queue** – Uses a separate mutex for the head and tail, with the queue operations protected by mutexes to ensure thread safety.
2. **Lock-Free Queue** – Implements a lock-free queue using atomic operations to ensure thread safety without using traditional locking mechanisms like mutexes.

## **How It Works**  

### **Michael and Scott's Queue (msqueue.go)**  
- **Head and Tail Locks:** This queue uses separate locks for the head and tail, allowing finer-grained locking while maintaining thread safety.
- **Enqueue Operation:** Locks the tail to insert a new node into the queue.
- **Dequeue Operation:** Locks the head to remove a node from the front of the queue.
- **Tradeoff:** Allows concurrent operations at the head and tail but requires locking for each operation, which can cause contention under high concurrency.

### **Lock-Free Queue (lfqueue.go)**  
- **No Locks:** This queue implements a lock-free design where atomic operations are used to ensure thread safety without requiring locks. It uses a spin-lock mechanism to handle concurrent access.
- **Enqueue Operation:** Adds elements to the queue using atomic operations to ensure safe concurrent insertion.
- **Dequeue Operation:** Removes elements using atomic operations, ensuring thread safety even when multiple threads are accessing the queue.
- **Tradeoff:** The lock-free design allows better performance under high concurrency but is more complex to implement.

## **Benchmarking (lfqueue_test.go and msqueue_test.go)**  
- The benchmarking tests are implemented in two files: 
  - **`lfqueue_test.go`** benchmarks the lock-free queue.
  - **`msqueue_test.go`** benchmarks Michael and Scott's queue.
  
- **Benchmarking Process:**
  - Both queues are benchmarked by enqueuing and dequeuing a set number of elements.
  - For **Michael and Scott's Queue**, the operations are measured with head and tail locking.
  - For **Lock-Free Queue**, the operations are measured with atomic operations.

- **Metrics:**
  - **Enqueue Time:** Measures the time taken to enqueue elements.
  - **Dequeue Time:** Measures the time taken to dequeue elements.

### Prerequisites

- [Go](https://golang.org/dl/) is installed to run the program.

### Cloning the Repository
Clone the repository with the following commands:

```bash
```bash
git clone https://github.com/MikeS3/OS_Homework.git
cd OS_Homework
cd HW4
cd homework4
```

### Running the Code
To run the code type this command in the terminal in the directory with the files for this project
```bash
go test -benchmem -run=^$ -bench ^BenchmarkQueue_Enqueue$ homework4/MichealScott; go test -benchmem -run=^$ -bench ^BenchmarkQueue_Enqueue$ homework4/LockFree
```

### Example Output:
goos: windows  
goarch: amd64  
pkg: homework4/LockFree  
cpu: AMD Ryzen 7 4700U with Radeon Graphics  
BenchmarkQueue_Enqueue-8         8480498               122.5 ns/op            31 B/op          1 allocs/op  
PASS  
ok      homework4/LockFree      2.463s  
