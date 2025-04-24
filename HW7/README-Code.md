# RAID Simulation in Go  
-format structure created by Claude 3.7 updated by us

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
go run go_run_visualization.go
```
## Project Structure

### RAID Interface
All RAID implementations adhere to this common interface:

```go
type RAID interface {
    Write(blockNum int, data []byte) error
    Read(blockNum int) ([]byte, error)
    Initialize() error
    CleanUp() error
    GetEffectiveCapacity() int
    GetName() string
}
```

### Disk Implementation
The project simulates physical disks using files:

```go
type Disk struct {
    file *os.File
    path string
    mu   sync.Mutex
}
```

Each disk provides basic read/write operations and proper synchronization (using `fsync()`) to simulate actual disk behavior.

### RAID Implementations

#### RAID-0 (Striping)
- Distributes data across all disks in a round-robin fashion
- No redundancy or fault tolerance
- Maximum performance and capacity utilization
- Main functions:
  - `Write()`: Stripes data across disks
  - `Read()`: Reads from appropriate disk based on block number

#### RAID-1 (Mirroring)
- Writes identical data to all disks
- Provides full redundancy with N copies
- Reads from the first disk (could be optimized for read balancing)
- Main functions:
  - `Write()`: Duplicates data to all disks
  - `Read()`: Reads from first disk

#### RAID-4 (Dedicated Parity)
- Stripes data across N-1 disks
- Dedicates one disk for parity
- Uses XOR for parity calculation
- Main functions:
  - `Write()`: Writes data to data disk and calculates/writes parity
  - `Read()`: Reads from appropriate data disk

#### RAID-5 (Distributed Parity)
- Stripes data across all disks
- Distributes parity information across all disks
- Uses XOR for parity calculation
- Main functions:
  - `Write()`: Writes data and calculates/distributes parity
  - `Read()`: Reads from appropriate data disk
  - `getParityDisk()`: Determines which disk stores parity for each strip

### Testing Framework

The project includes comprehensive tests for all RAID implementations:

1. **Basic Functionality Tests**
   - Tests write/read operations for all RAID levels
   - Verifies data integrity across operations

2. **Parity Calculation Test**
   - Validates XOR parity calculation
   - Confirms data recovery using parity

3. **Capacity Calculation Test**
   - Verifies correct capacity calculations for each RAID level

4. **RAID-5 Parity Rotation Test**
   - Tests the parity rotation logic in RAID-5

### Benchmarking

The code includes a benchmarking system that:
1. Writes a specified amount of data (100MB by default)
2. Reads the data back
3. Measures and reports performance metrics
4. Visualizes results using ASCII charts

## Constants and Configuration

```go
const (
    BlockSize  = 4096      // 4KB blocks
    NumDisks   = 5         // 5 disk simulation
    NumBlocks  = 10000     // Total logical blocks
    DataSize   = 100 * 1024 * 1024 // 100MB for benchmarking
)
```

## Running Tests

Use the following command to run all tests:

```bash
go run go_run_visualization.go
```

## Implementation Notes

1. **Disk Synchronization**
   - Each disk operation uses mutex locks to ensure thread safety
   - `fsync()` is called after each write to simulate physical disk delays

2. **Error Handling**
   - The implementation handles basic I/O errors
   - Edge cases like disk failures are not handled (as per assignment)

3. **Parity Calculation**
   - Uses the XOR operation for parity calculations
   - For RAID-4, parity is stored on a dedicated disk
   - For RAID-5, parity is distributed across all disks

4. **Performance Considerations**
   - The implementation prioritizes correctness over performance
   - Real-world optimizations like caching are not implemented
   - Synchronous operations are used to better simulate real disk behavior
