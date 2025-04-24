package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

const (
	BlockSize  = 4096 // 4KB
	NumDisks   = 5
	NumBlocks  = 10000 // Total logical blocks
	DataSize   = 100 * 1024 * 1024 // 100MB for benchmarking
)

// RAID interface as specified in the assignment
type RAID interface {
	Write(blockNum int, data []byte) error
	Read(blockNum int) ([]byte, error)
	Initialize() error
	CleanUp() error
	GetEffectiveCapacity() int
	GetName() string
}

// Disk represents a simulated physical disk
type Disk struct {
	file *os.File
	path string
	mu   sync.Mutex
}

// NewDisk creates a new simulated disk
func NewDisk(path string) (*Disk, error) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}

	return &Disk{
		file: file,
		path: path,
	}, nil
}

// Read reads a block from the disk
func (d *Disk) Read(blockNum int, buffer []byte) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	offset := int64(blockNum) * int64(BlockSize)
	_, err := d.file.Seek(offset, io.SeekStart)
	if err != nil {
		return err
	}

	_, err = io.ReadFull(d.file, buffer)
	if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
		return err
	}

	return nil
}

// Write writes a block to the disk
func (d *Disk) Write(blockNum int, data []byte) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	offset := int64(blockNum) * int64(BlockSize)
	_, err := d.file.Seek(offset, io.SeekStart)
	if err != nil {
		return err
	}

	_, err = d.file.Write(data)
	if err != nil {
		return err
	}

	// Ensure data is flushed to simulate real disk write delays
	return d.file.Sync()
}

// Close closes the disk
func (d *Disk) Close() error {
	return d.file.Close()
}

// Delete deletes the disk file
func (d *Disk) Delete() error {
	err := d.file.Close()
	if err != nil {
		return err
	}
	return os.Remove(d.path)
}

// RAID0 implements striping across disks
type RAID0 struct {
	disks      []*Disk
	blockSize  int
	numDisks   int
}

func NewRAID0() *RAID0 {
	return &RAID0{
		blockSize: BlockSize,
		numDisks:  NumDisks,
	}
}

func (r *RAID0) GetName() string {
	return "RAID0"
}

func (r *RAID0) Initialize() error {
	r.disks = make([]*Disk, r.numDisks)
	for i := 0; i < r.numDisks; i++ {
		disk, err := NewDisk(fmt.Sprintf("disk%d.dat", i))
		if err != nil {
			return err
		}
		r.disks[i] = disk
	}
	return nil
}

func (r *RAID0) CleanUp() error {
	for _, disk := range r.disks {
		err := disk.Delete()
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *RAID0) GetEffectiveCapacity() int {
	return r.numDisks * NumBlocks
}

func (r *RAID0) Write(blockNum int, data []byte) error {
	if len(data) != r.blockSize {
		return errors.New("data size does not match block size")
	}

	diskNum := blockNum % r.numDisks
	diskBlockNum := blockNum / r.numDisks

	return r.disks[diskNum].Write(diskBlockNum, data)
}

func (r *RAID0) Read(blockNum int) ([]byte, error) {
	diskNum := blockNum % r.numDisks
	diskBlockNum := blockNum / r.numDisks

	data := make([]byte, r.blockSize)
	err := r.disks[diskNum].Read(diskBlockNum, data)
	return data, err
}

// RAID1 implements mirroring across disks
type RAID1 struct {
	disks      []*Disk
	blockSize  int
	numDisks   int
}

func NewRAID1() *RAID1 {
	return &RAID1{
		blockSize: BlockSize,
		numDisks:  NumDisks,
	}
}

func (r *RAID1) GetName() string {
	return "RAID1"
}

func (r *RAID1) Initialize() error {
	r.disks = make([]*Disk, r.numDisks)
	for i := 0; i < r.numDisks; i++ {
		disk, err := NewDisk(fmt.Sprintf("disk%d.dat", i))
		if err != nil {
			return err
		}
		r.disks[i] = disk
	}
	return nil
}

func (r *RAID1) CleanUp() error {
	for _, disk := range r.disks {
		err := disk.Delete()
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *RAID1) GetEffectiveCapacity() int {
	return NumBlocks
}

func (r *RAID1) Write(blockNum int, data []byte) error {
	if len(data) != r.blockSize {
		return errors.New("data size does not match block size")
	}

	// Write to all disks for mirroring
	for _, disk := range r.disks {
		err := disk.Write(blockNum, data)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *RAID1) Read(blockNum int) ([]byte, error) {
	// Read from the first disk (could implement read balancing here)
	data := make([]byte, r.blockSize)
	err := r.disks[0].Read(blockNum, data)
	return data, err
}

// RAID4 implements block-level striping with a dedicated parity disk
type RAID4 struct {
	disks      []*Disk
	blockSize  int
	numDisks   int
	parityDisk int
	dataDisks  int
}

func NewRAID4() *RAID4 {
	return &RAID4{
		blockSize:  BlockSize,
		numDisks:   NumDisks,
		parityDisk: NumDisks - 1, // Last disk is parity
		dataDisks:  NumDisks - 1, // All except parity disk
	}
}

func (r *RAID4) GetName() string {
	return "RAID4"
}

func (r *RAID4) Initialize() error {
	r.disks = make([]*Disk, r.numDisks)
	for i := 0; i < r.numDisks; i++ {
		disk, err := NewDisk(fmt.Sprintf("disk%d.dat", i))
		if err != nil {
			return err
		}
		r.disks[i] = disk
	}
	return nil
}

func (r *RAID4) CleanUp() error {
	for _, disk := range r.disks {
		err := disk.Delete()
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *RAID4) GetEffectiveCapacity() int {
	return r.dataDisks * NumBlocks
}

func (r *RAID4) Write(blockNum int, data []byte) error {
	if len(data) != r.blockSize {
		return errors.New("data size does not match block size")
	}

	stripNum := blockNum / r.dataDisks
	diskNum := blockNum % r.dataDisks

	// Write data to data disk
	err := r.disks[diskNum].Write(stripNum, data)
	if err != nil {
		return err
	}

	// Calculate and write parity
	parity := make([]byte, r.blockSize)
	
	// Read data from all disks for this strip to calculate parity
	for i := 0; i < r.dataDisks; i++ {
		if i == diskNum {
			// We already have this data, XOR it with parity
			for j := 0; j < r.blockSize; j++ {
				parity[j] ^= data[j]
			}
		} else {
			// Read data from other disks
			blockData := make([]byte, r.blockSize)
			err := r.disks[i].Read(stripNum, blockData)
			if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
				return err
			}
			
			// XOR with parity
			for j := 0; j < r.blockSize; j++ {
				parity[j] ^= blockData[j]
			}
		}
	}

	// Write parity to parity disk
	return r.disks[r.parityDisk].Write(stripNum, parity)
}

func (r *RAID4) Read(blockNum int) ([]byte, error) {
	stripNum := blockNum / r.dataDisks
	diskNum := blockNum % r.dataDisks

	data := make([]byte, r.blockSize)
	err := r.disks[diskNum].Read(stripNum, data)
	return data, err
}

// RAID5 implements block-level striping with distributed parity
type RAID5 struct {
	disks     []*Disk
	blockSize int
	numDisks  int
	dataDisks int
}

func NewRAID5() *RAID5 {
	return &RAID5{
		blockSize: BlockSize,
		numDisks:  NumDisks,
		dataDisks: NumDisks - 1, // One disk's worth of capacity is used for parity
	}
}

func (r *RAID5) GetName() string {
	return "RAID5"
}

func (r *RAID5) Initialize() error {
	r.disks = make([]*Disk, r.numDisks)
	for i := 0; i < r.numDisks; i++ {
		disk, err := NewDisk(fmt.Sprintf("disk%d.dat", i))
		if err != nil {
			return err
		}
		r.disks[i] = disk
	}
	return nil
}

func (r *RAID5) CleanUp() error {
	for _, disk := range r.disks {
		err := disk.Delete()
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *RAID5) GetEffectiveCapacity() int {
	return r.dataDisks * NumBlocks
}

// getParityDisk returns the disk number that stores parity for a given strip
func (r *RAID5) getParityDisk(stripNum int) int {
	// Rotating parity: parity for strip N goes to disk (N % numDisks)
	return (r.numDisks - 1 - (stripNum % r.numDisks))
}

func (r *RAID5) Write(blockNum int, data []byte) error {
	if len(data) != r.blockSize {
		return errors.New("data size does not match block size")
	}

	stripNum := blockNum / r.dataDisks
	stripOffset := blockNum % r.dataDisks
	
	parityDisk := r.getParityDisk(stripNum)
	
	// Adjust the disk number if it's affected by the parity placement
	diskNum := stripOffset
	if diskNum >= parityDisk {
		diskNum++
	}
	
	// If diskNum is now equal to numDisks, wrap around
	if diskNum >= r.numDisks {
		diskNum = 0
	}

	// Write data to the data disk
	err := r.disks[diskNum].Write(stripNum, data)
	if err != nil {
		return err
	}

	// Calculate and write parity
	parity := make([]byte, r.blockSize)
	
	// Read data from all disks for this strip to calculate parity
	for i := 0; i < r.numDisks; i++ {
		if i == parityDisk {
			continue // Skip parity disk when calculating
		}
		
		if i == diskNum {
			// We already have this data, XOR it with parity
			for j := 0; j < r.blockSize; j++ {
				parity[j] ^= data[j]
			}
		} else {
			// Read data from other data disks
			blockData := make([]byte, r.blockSize)
			err := r.disks[i].Read(stripNum, blockData)
			if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
				return err
			}
			
			// XOR with parity
			for j := 0; j < r.blockSize; j++ {
				parity[j] ^= blockData[j]
			}
		}
	}

	// Write parity to parity disk
	return r.disks[parityDisk].Write(stripNum, parity)
}

func (r *RAID5) Read(blockNum int) ([]byte, error) {
	stripNum := blockNum / r.dataDisks
	stripOffset := blockNum % r.dataDisks
	
	parityDisk := r.getParityDisk(stripNum)
	
	// Adjust the disk number if it's affected by the parity placement
	diskNum := stripOffset
	if diskNum >= parityDisk {
		diskNum++
	}
	
	// If diskNum is now equal to numDisks, wrap around
	if diskNum >= r.numDisks {
		diskNum = 0
	}

	data := make([]byte, r.blockSize)
	err := r.disks[diskNum].Read(stripNum, data)
	return data, err
}

// RunBenchmark runs benchmark tests on a RAID implementation
func RunBenchmark(raid RAID, numBlocks int) (writeTime, readTime time.Duration, err error) {
	// Initialize RAID
	err = raid.Initialize()
	if err != nil {
		return 0, 0, err
	}
	defer raid.CleanUp()

	// Generate test data
	testData := make([]byte, BlockSize)
	for i := range testData {
		testData[i] = byte(i % 256)
	}

	// Measure write performance
	writeStart := time.Now()
	for i := 0; i < numBlocks; i++ {
		err = raid.Write(i, testData)
		if err != nil {
			return 0, 0, err
		}
	}
	writeTime = time.Since(writeStart)

	// Measure read performance
	readStart := time.Now()
	for i := 0; i < numBlocks; i++ {
		_, err = raid.Read(i)
		if err != nil {
			return 0, 0, err
		}
	}
	readTime = time.Since(readStart)

	return writeTime, readTime, nil
}

// FormatDuration formats a duration as seconds with 2 decimal places
func FormatDuration(d time.Duration) string {
	return fmt.Sprintf("%.2f seconds", d.Seconds())
}

// CalculateSpeed calculates the speed in MB/s
func CalculateSpeed(dataSize int, duration time.Duration) float64 {
	return float64(dataSize) / 1024 / 1024 / duration.Seconds()
}

func main() {
	// Calculate how many blocks for 100MB
	numBenchmarkBlocks := DataSize / BlockSize
	if DataSize%BlockSize != 0 {
		numBenchmarkBlocks++
	}

	fmt.Printf("RAID Simulation Benchmark\n")
	fmt.Printf("Block Size: %d bytes\n", BlockSize)
	fmt.Printf("Number of Disks: %d\n", NumDisks)
	fmt.Printf("Data Size: %d MB\n", DataSize/1024/1024)
	fmt.Printf("Number of Blocks: %d\n\n", numBenchmarkBlocks)

	// Run benchmarks for each RAID level
	raids := []RAID{
		NewRAID0(),
		NewRAID1(),
		NewRAID4(),
		NewRAID5(),
	}

	fmt.Printf("%-8s %-15s %-15s %-15s %-15s %-15s %-15s\n",
		"RAID", "Write Time", "Write Speed", "Read Time", "Read Speed", "Effective Cap", "Overhead")

	for _, raid := range raids {
		writeTime, readTime, err := RunBenchmark(raid, numBenchmarkBlocks)
		if err != nil {
			log.Fatalf("Error running benchmark for %s: %v", raid.GetName(), err)
		}

		writeSpeed := CalculateSpeed(DataSize, writeTime)
		readSpeed := CalculateSpeed(DataSize, readTime)
		effectiveCap := raid.GetEffectiveCapacity() * BlockSize / (1024 * 1024) // in MB
		overhead := 100.0 - (float64(effectiveCap) / float64(NumDisks*NumBlocks*BlockSize/(1024*1024)) * 100.0)

		fmt.Printf("%-8s %-15s %-15.2f %-15s %-15.2f %-15d %-15.2f\n",
			raid.GetName(),
			FormatDuration(writeTime),
			writeSpeed,
			FormatDuration(readTime),
			readSpeed,
			effectiveCap,
			overhead)
	}

	// Analysis and comparison with textbook expectations
	fmt.Printf("\nAnalysis:\n")
	fmt.Printf("- RAID0: Provides the highest performance but no redundancy.\n")
	fmt.Printf("- RAID1: Provides redundancy but at the cost of capacity (1/N of total capacity).\n")
	fmt.Printf("- RAID4: Better capacity utilization than RAID1 but parity disk becomes a bottleneck for writes.\n")
	fmt.Printf("- RAID5: Distributes parity to avoid the bottleneck in RAID4 while maintaining redundancy.\n")
	fmt.Printf("\nIf the performance trends match textbook expectations, RAID0 should be fastest for both reads and writes,\n")
	fmt.Printf("while RAID5 should offer better write performance than RAID4 due to distributed parity.\n")
}