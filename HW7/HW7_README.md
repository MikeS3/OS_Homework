# HW7 RAID Simulation in Go
Comparing the execution times of a global lock linked list and a hand-over-hand locking linked list 

## **Program Overview**  
This program implements and benchmarks two types of concurrent linked lists: 
1. **Global Lock Linked List** – Uses a single mutex to protect the entire list. 
2. **Hand-Over-Hand Locking Linked List** – Uses one mutex per node, allowing finer-grained locking. 

## **How It Works**  

### **Global Lock Linked List (global_lock_list.go)**  
- A **single global mutex** (sync.Mutex) protects the entire list. 
- **Insert Operation:** Locks the entire list, checks if the key exists, then inserts if the key isnt found. 
- **Lookup Operation:** Locks the entire list, searches for the key, then unlocks. 
- **Tradeoff:** Simple, but can cause high contention when multiple threads are trying to access the list. 

### **Hand-Over-Hand Locking Linked List (hand_over_hand_list.go)**  
- Each node has its own lock, which allows multiple threads to traverse different parts of the list at the same time 
- **Insert Operation:** Locks nodes one by one while also searching for duplicate nodes, then inserts at the head. 
- **Lookup Operation:** Uses hand-over-hand locking which locks the next node before unlocking the current node. 
- **Tradeoff:** Allows more concurrent access, but it has higher locking overhead. 

## **Benchmarking (main.go)**  
- Spawns **10 concurrent threads**, each performing **1,000 insert or lookup operations**. 
- **Measures execution time** for: 
  - **Inserts (`benchmark(list, true)`)** 
  - **Lookups (`benchmark(list, false)`)** 
- Runs benchmarks on both the hand-over-hand and global locking implementations and then prints the time for each implementation. 

### Prerequisites

- [Go](https://golang.org/dl/) is installed so you can run program.

### Cloning the Repository
Clone the repository with the following commands:

```bash
git clone https://github.com/MikeS3/OS_Homework.git
cd OS_Homework
cd HW3
```

### Running the Code
To run the code type this command in the terminal in the directory with the files for this project
```bash
go run main.go global_lock_list.go hand_over_hand_list.go
```

## Results
The expected output will show the execution times for insert and lookup operations on both implementations of the queue:

Benchmarking Global Lock List
Time elapsed: 7.555701ms | Successful operations: 1000
Time elapsed: 6.9879ms | Successful operations: 10000
Benchmarking Hand-over-Hand Lock List
Time elapsed: 66.5193ms | Successful operations: 1000
Time elapsed: 60.303701ms | Successful operations: 10000
