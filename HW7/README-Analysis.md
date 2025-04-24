# RAID Simulation in Go - Performance Analysis
-Format created by Claude 3.7 updated by us

## Overview
This document provides an analysis of the performance characteristics of different RAID levels (0, 1, 4, 5) as implemented in our Go-based simulation. It compares theoretical expectations with measured results and offers insights into the trade-offs between performance, capacity, and reliability.

## Benchmark Results

Our simulation produced the following benchmark results:

```
RAID Simulation Benchmark
Block Size: 4096 bytes
Number of Disks: 5
Data Size: 100 MB
Number of Blocks: 25600

RAID     Write Time      Write Speed     Read Time       Read Speed      Effective Cap   Overhead
RAID0    140.44 seconds  0.71 MB/s       0.08 seconds    1193.27 MB/s    195 MB          0.00%
RAID1    747.39 seconds  0.13 MB/s       0.08 seconds    1220.81 MB/s    39 MB           80.00%
RAID4    264.37 seconds  0.38 MB/s       0.06 seconds    1591.71 MB/s    156 MB          20.00%
RAID5    254.39 seconds  0.39 MB/s       0.06 seconds    1696.46 MB/s    156 MB          20.00%
```

## Theoretical vs. Measured Performance

### RAID-0 (Striping)

**Theoretical Model**: N · S (where N = number of disks, S = single disk speed)
- **Read**: N · S for sequential reads, N · R for random reads
- **Write**: N · S for sequential writes, N · R for random writes
- **Capacity**: N · B (where B = capacity of one disk)
- **Reliability**: 0 (failure of any disk results in data loss)

**Measured Results**:
- **Read**: 1193.27 MB/s
- **Write**: 0.71 MB/s
- **Capacity**: 195 MB (using 5 disks)

**Analysis**: Our measured sequential read throughput aligns with theoretical expectations, showing the benefit of striping. Write performance is lower than expected due to our use of `fsync()` to simulate real disk writing. The capacity utilization of 100% matches theoretical predictions.

### RAID-1 (Mirroring)

**Theoretical Model**: 
- **Read**: (N/2) · S for sequential reads, N · R for random reads
- **Write**: (N/2) · S for sequential writes, (N/2) · R for random writes
- **Capacity**: (N · B)/2 (traditional mirror pair setup)
- **Reliability**: 1 (can survive one disk failure for sure)

**Measured Results**:
- **Read**: 1220.81 MB/s
- **Write**: 0.13 MB/s
- **Capacity**: 39 MB (using 5 disks)

**Analysis**: Our read performance exceeds theoretical expectations because we read from only one disk rather than balancing across mirrors. Write performance is significantly lower than RAID-0 due to writing identical data to all five disks. Our capacity is 1/5 of total capacity because we mirror to all disks rather than using disk pairs.

### RAID-4 (Dedicated Parity)

**Theoretical Model**:
- **Read**: (N-1) · S for sequential reads, (N-1) · R for random reads
- **Write**: (N-1) · S for sequential writes, 1/2 · R for random writes (parity bottleneck)
- **Capacity**: (N-1) · B
- **Reliability**: 1 (can survive one disk failure)

**Measured Results**:
- **Read**: 1591.71 MB/s
- **Write**: 0.38 MB/s
- **Capacity**: 156 MB (using 5 disks)

**Analysis**: Read performance matches theoretical expectations. Write performance is affected by the parity disk bottleneck but still outperforms RAID-1 due to the reduced redundancy overhead. Capacity utilization of 80% (4/5 disks) aligns with theoretical expectations.

### RAID-5 (Distributed Parity)

**Theoretical Model**:
- **Read**: (N-1) · S for sequential reads, N · R for random reads
- **Write**: (N-1) · S for sequential writes, 1/4 · R for random writes
- **Capacity**: (N-1) · B
- **Reliability**: 1 (can survive one disk failure)

**Measured Results**:
- **Read**: 1696.46 MB/s
- **Write**: 0.39 MB/s
- **Capacity**: 156 MB (using 5 disks)

**Analysis**: Our RAID-5 implementation shows slightly better read and write performance than RAID-4, which aligns with theoretical expectations due to the distribution of parity across disks, eliminating the parity disk bottleneck. Capacity utilization matches RAID-4 at 80% (4/5 disks).

## Performance Visualizations

```
Write Speed (MB/s):
RAID0  [██████████████████████████████████████████████████] 0.71
RAID1  [████████] 0.13
RAID4  [██████████████████████████] 0.38
RAID5  [███████████████████████████] 0.39

Read Speed (MB/s):
RAID0  [███████████████████████████████████████] 1193.27
RAID1  [████████████████████████████████████████] 1220.81
RAID4  [██████████████████████████████████████████████████] 1591.71
RAID5  [██████████████████████████████████████████████████████] 1696.46

Effective Capacity (MB):
RAID0  [██████████████████████████████████████████████████] 195.00
RAID1  [██████████] 39.00
RAID4  [████████████████████████████████████] 156.00
RAID5  [████████████████████████████████████] 156.00

Storage Overhead (%):
RAID0  [                                                  ] 0.00
RAID1  [██████████████████████████████████████████████████] 80.00
RAID4  [███████████] 20.00
RAID5  [███████████] 20.00
```

## Observations vs. Textbook Expectations

1. **Read Performance**:
   - Textbook predicts RAID-0 and RAID-5 should have similar read performances
   - Our simulation shows RAID-5 outperforming RAID-0 for reads
   - This difference may be due to our implementation details and caching effects

2. **Write Performance**:
   - Textbook suggests RAID-4 should have significantly worse write performance than RAID-5
   - Our simulation shows only a marginal difference (0.38 vs 0.39 MB/s)
   - This may be because our implementation still requires reading all blocks to calculate parity in both cases

3. **Capacity Utilization**:
   - Our results match theoretical expectations perfectly:
     - RAID-0: 100% utilization
     - RAID-1: 20% utilization (with our full 5-disk mirroring)
     - RAID-4/RAID-5: 80% utilization

4. **Latency**:
   - The textbook suggests RAID-4 and RAID-5 should have twice the write latency of RAID-0
   - Our simulation shows the write time for RAID-4/5 is about 1.9x that of RAID-0 (264/254 vs 140 seconds)
   - This closely matches theoretical expectations

## Factors Affecting Simulation Results

1. **File System Overhead**:
   - Our simulation uses the host file system, which adds overhead
   - Real RAID systems often use raw block devices

2. **Synchronization**:
   - We use `fsync()` to simulate disk latency, which is more costly than real hardware would be
   - This explains the generally low write speeds across all RAID levels

3. **Sequential vs. Random Access**:
   - Our benchmarks focus on sequential access
   - Real-world workloads often include random access patterns
   - According to the textbook, RAID-5 has an advantage in random reads over RAID-4

4. **Implementation Specifics**:
   - Our RAID-1 mirrors to all disks rather than using disk pairs
   - Our parity calculation reads all blocks rather than using more efficient algorithms

## Real-World Implications

When selecting a RAID level for production use, consider:

1. **Application Read/Write Patterns**:
   - Read-intensive workloads: RAID-5 provides excellent performance with redundancy
   - Write-intensive workloads: RAID-0 provides best performance, but with no redundancy

2. **Data Importance**:
   - Critical data: RAID-1 or RAID-5 for redundancy
   - Temporary/Reproducible data: RAID-0 for maximum performance

3. **Budget Constraints**:
   - High budget: RAID-1 provides simplest redundancy but at high cost
   - Limited budget: RAID-5 provides good balance of redundancy and capacity

4. **Performance Requirements**:
   - Maximum read speed: RAID-5 showed best results in our tests
   - Maximum write speed: RAID-0 is significantly faster

## Conclusion

Our simulation confirms the fundamental characteristics of different RAID levels described in the textbook. The performance trends generally align with theoretical expectations, with some variations due to implementation specifics.

The results validate the common use cases:
- RAID-0 for maximum performance and capacity
- RAID-1 for maximum reliability
- RAID-5 for balanced performance, capacity, and reliability

This simulation provides a practical demonstration of RAID concepts and performance characteristics, helping to understand the trade-offs involved in selecting an appropriate RAID level for different applications.
