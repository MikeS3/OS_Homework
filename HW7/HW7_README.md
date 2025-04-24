# RAID Simulation in Go

## Overview
This repository simulates different RAID levels (RAID 0, 1, 4, and 5) in Go. It aims to evaluate the write and read performance of each RAID level, compare their effective capacities, and understand the trade-offs in terms of redundancy, speed, and capacity utilization.

## Features
- RAID Levels: 0, 1, 4, 5
- Performance benchmarking (write/read speed)
- Effective capacity and overhead calculations
- Parity calculations (for RAID 4 and RAID 5)
- Test cases to validate data integrity and RAID functionality

## Prerequisites
Before running the simulation and benchmarks, ensure you have the following installed:

1. **Install Go**:  
   Follow the instructions at [https://golang.org/doc/install](https://golang.org/doc/install) to install Go.

2. **Clone the Repository**:  
   Clone the repository using the following commands:

```bash
git clone https://github.com/MikeS3/OS_Homework.git
cd OS_Homework
cd HW7
```

## How to Run the Simulation

### Running the Benchmark:
To run the RAID simulation and benchmarks, execute the following command:
```bash
go run main.go
```

## Code Walkthrough
The main functionality is broken down into several components:

### RAID Interface:
The RAID interface defines common methods that all RAID levels (RAID 0, 1, 4, 5) should implement, such as Write, Read, Initialize, CleanUp, and GetEffectiveCapacity.

### RAID Implementations:
Each RAID level is implemented as a struct with specific logic for data distribution and parity handling.
- **RAID0** - Stripe-only (no redundancy)
- **RAID1** - Mirroring (redundancy)
- **RAID4** - Parity (single parity disk)
- **RAID5** - Parity (distributed parity across disks)

### Benchmarking:
The benchmark measures the time taken for write and read operations and calculates the throughput (MB/s). It compares the performance of each RAID level by writing and reading a defined amount of data across blocks.

### Effective Capacity Calculation:
Each RAID level calculates its effective storage capacity based on the number of disks, redundancy type, and parity distribution.

### Parity Calculation:
Parity is calculated using the XOR operation for RAID 4 and RAID 5, ensuring data can be reconstructed in case of a disk failure.

## Running the Benchmark and Analysis
To run the benchmarks, use the following command:
```bash
go run main.go
```

## Benchmark Results
When the benchmarks are executed, the output will be similar to this:
```
RAID Simulation Benchmark
Block Size: 4096 bytes
Number of Disks: 5
Data Size: 100 MB
Number of Blocks: 25600

RAID     Write Time      Write Speed     Read Time       Read Speed      Effective Cap   Overhead
RAID0    140.44 seconds  0.71            0.08 seconds    1193.27         195             0.00
RAID1    747.39 seconds  0.13            0.08 seconds    1220.81         39              80.00
RAID4    264.37 seconds  0.38            0.06 seconds    1591.71         156             20.00
RAID5    254.39 seconds  0.39            0.06 seconds    1696.46         156             20.00
```

### Analysis:
- **RAID0**: Provides the highest performance but no redundancy.
- **RAID1**: Provides redundancy but at the cost of capacity (1/N of total capacity).
- **RAID4**: Better capacity utilization than RAID1 but parity disk becomes a bottleneck for writes.
- **RAID5**: Distributes parity to avoid the bottleneck in RAID4 while maintaining redundancy.

If the performance trends match textbook expectations, RAID0 should be fastest for both reads and writes, while RAID5 should offer better write performance than RAID4 due to distributed parity.
