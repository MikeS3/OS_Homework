package main

import (
	"bytes"
	"testing"
)

const TestBlockSize = 4096

// TestRAID0 tests basic RAID0 functionality
func TestRAID0(t *testing.T) {
	raid := NewRAID0()
	raid.blockSize = TestBlockSize
	raid.numDisks = 5

	err := raid.Initialize()
	if err != nil {
		t.Fatalf("Failed to initialize RAID0: %v", err)
	}
	defer raid.CleanUp()

	// Test data
	testData := make([]byte, TestBlockSize)
	for i := range testData {
		testData[i] = byte(i % 256)
	}

	// Test across multiple blocks
	for blockNum := 0; blockNum < 20; blockNum++ {
		// Write data
		err = raid.Write(blockNum, testData)
		if err != nil {
			t.Fatalf("Failed to write to RAID0 block %d: %v", blockNum, err)
		}

		// Read data back
		readData, err := raid.Read(blockNum)
		if err != nil {
			t.Fatalf("Failed to read from RAID0 block %d: %v", blockNum, err)
		}

		// Verify data
		if !bytes.Equal(readData, testData) {
			t.Errorf("Data mismatch in RAID0 block %d", blockNum)
		}
	}
}

// TestRAID1 tests basic RAID1 functionality
func TestRAID1(t *testing.T) {
	raid := NewRAID1()
	raid.blockSize = TestBlockSize
	raid.numDisks = 5

	err := raid.Initialize()
	if err != nil {
		t.Fatalf("Failed to initialize RAID1: %v", err)
	}
	defer raid.CleanUp()

	// Test data
	testData := make([]byte, TestBlockSize)
	for i := range testData {
		testData[i] = byte(i % 256)
	}

	// Test across multiple blocks
	for blockNum := 0; blockNum < 20; blockNum++ {
		// Write data
		err = raid.Write(blockNum, testData)
		if err != nil {
			t.Fatalf("Failed to write to RAID1 block %d: %v", blockNum, err)
		}

		// Read data back
		readData, err := raid.Read(blockNum)
		if err != nil {
			t.Fatalf("Failed to read from RAID1 block %d: %v", blockNum, err)
		}

		// Verify data
		if !bytes.Equal(readData, testData) {
			t.Errorf("Data mismatch in RAID1 block %d", blockNum)
		}
	}
}

// TestRAID4 tests basic RAID4 functionality
func TestRAID4(t *testing.T) {
	raid := NewRAID4()
	raid.blockSize = TestBlockSize
	raid.numDisks = 5
	raid.parityDisk = 4
	raid.dataDisks = 4

	err := raid.Initialize()
	if err != nil {
		t.Fatalf("Failed to initialize RAID4: %v", err)
	}
	defer raid.CleanUp()

	// Test data
	testData := make([]byte, TestBlockSize)
	for i := range testData {
		testData[i] = byte(i % 256)
	}

	// Test across multiple blocks
	for blockNum := 0; blockNum < 20; blockNum++ {
		// Write data
		err = raid.Write(blockNum, testData)
		if err != nil {
			t.Fatalf("Failed to write to RAID4 block %d: %v", blockNum, err)
		}

		// Read data back
		readData, err := raid.Read(blockNum)
		if err != nil {
			t.Fatalf("Failed to read from RAID4 block %d: %v", blockNum, err)
		}

		// Verify data
		if !bytes.Equal(readData, testData) {
			t.Errorf("Data mismatch in RAID4 block %d", blockNum)
		}
	}
}

// TestRAID5 tests basic RAID5 functionality
func TestRAID5(t *testing.T) {
	raid := NewRAID5()
	raid.blockSize = TestBlockSize
	raid.numDisks = 5
	raid.dataDisks = 4

	err := raid.Initialize()
	if err != nil {
		t.Fatalf("Failed to initialize RAID5: %v", err)
	}
	defer raid.CleanUp()

	// Test data
	testData := make([]byte, TestBlockSize)
	for i := range testData {
		testData[i] = byte(i % 256)
	}

	// Test across multiple blocks
	for blockNum := 0; blockNum < 20; blockNum++ {
		// Write data
		err = raid.Write(blockNum, testData)
		if err != nil {
			t.Fatalf("Failed to write to RAID5 block %d: %v", blockNum, err)
		}

		// Read data back
		readData, err := raid.Read(blockNum)
		if err != nil {
			t.Fatalf("Failed to read from RAID5 block %d: %v", blockNum, err)
		}

		// Verify data
		if !bytes.Equal(readData, testData) {
			t.Errorf("Data mismatch in RAID5 block %d", blockNum)
		}
	}
}

// TestParityCalculation tests XOR parity calculation
func TestParityCalculation(t *testing.T) {
	// Create test data
	data1 := []byte{0x01, 0x02, 0x03, 0x04}
	data2 := []byte{0x05, 0x06, 0x07, 0x08}
	data3 := []byte{0x09, 0x0A, 0x0B, 0x0C}

	// Calculate parity
	parity := make([]byte, 4)
	for i := 0; i < 4; i++ {
		parity[i] = data1[i] ^ data2[i] ^ data3[i]
	}

	// Now we should be able to recover any one of the data blocks using the others and parity
	// Recover data1
	recovered1 := make([]byte, 4)
	for i := 0; i < 4; i++ {
		recovered1[i] = data2[i] ^ data3[i] ^ parity[i]
	}

	// Recover data2
	recovered2 := make([]byte, 4)
	for i := 0; i < 4; i++ {
		recovered2[i] = data1[i] ^ data3[i] ^ parity[i]
	}

	// Recover data3
	recovered3 := make([]byte, 4)
	for i := 0; i < 4; i++ {
		recovered3[i] = data1[i] ^ data2[i] ^ parity[i]
	}

	// Verify recovery
	if !bytes.Equal(data1, recovered1) {
		t.Errorf("Failed to recover data1 using parity")
	}
	if !bytes.Equal(data2, recovered2) {
		t.Errorf("Failed to recover data2 using parity")
	}
	if !bytes.Equal(data3, recovered3) {
		t.Errorf("Failed to recover data3 using parity")
	}
}

// TestEffectiveCapacity tests the effective capacity calculation
func TestEffectiveCapacity(t *testing.T) {
	// RAID0: All capacity is usable
	raid0 := NewRAID0()
	raid0.numDisks = 5
	expected0 := 5 * NumBlocks
	if raid0.GetEffectiveCapacity() != expected0 {
		t.Errorf("RAID0 effective capacity incorrect: got %d, expected %d", 
			raid0.GetEffectiveCapacity(), expected0)
	}

	// RAID1: Only 1/N capacity is usable
	raid1 := NewRAID1()
	raid1.numDisks = 5
	expected1 := NumBlocks
	if raid1.GetEffectiveCapacity() != expected1 {
		t.Errorf("RAID1 effective capacity incorrect: got %d, expected %d", 
			raid1.GetEffectiveCapacity(), expected1)
	}

	// RAID4 and RAID5: (N-1)/N capacity is usable
	raid4 := NewRAID4()
	raid4.numDisks = 5
	raid4.dataDisks = 4
	expected4 := 4 * NumBlocks
	if raid4.GetEffectiveCapacity() != expected4 {
		t.Errorf("RAID4 effective capacity incorrect: got %d, expected %d", 
			raid4.GetEffectiveCapacity(), expected4)
	}

	raid5 := NewRAID5()
	raid5.numDisks = 5
	raid5.dataDisks = 4
	expected5 := 4 * NumBlocks
	if raid5.GetEffectiveCapacity() != expected5 {
		t.Errorf("RAID5 effective capacity incorrect: got %d, expected %d", 
			raid5.GetEffectiveCapacity(), expected5)
	}
}

// TestRAID5ParityRotation tests the parity rotation logic in RAID5
func TestRAID5ParityRotation(t *testing.T) {
	raid := NewRAID5()
	raid.numDisks = 5
	
	// Test that parity disk rotates correctly
	expectedParityDisk := []int{
		4, // Strip 0: Parity on disk 4
		3, // Strip 1: Parity on disk 3
		2, // Strip 2: Parity on disk 2
		1, // Strip 3: Parity on disk 1
		0, // Strip 4: Parity on disk 0
		4, // Strip 5: Back to disk 4
	}
	
	for strip := 0; strip < len(expectedParityDisk); strip++ {
		parityDisk := raid.getParityDisk(strip)
		if parityDisk != expectedParityDisk[strip] {
			t.Errorf("RAID5 parity rotation error for strip %d: got disk %d, expected disk %d", 
				strip, parityDisk, expectedParityDisk[strip])
		}
	}
}