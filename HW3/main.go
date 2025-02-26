package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numOperations = 1000 // Number of insert or lookup operations per thread
	numThreads    = 10   // Number of concurrent threads
)

// benchmark performs concurrent insertions or lookups on the given list
func benchmark(list interface{}, insert bool) {
	var wg sync.WaitGroup
	successCount := 0
	start := time.Now() // Start measuring time

	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			localSuccess := 0 // Track successful operations per thread

			for j := 0; j < numOperations; j++ {
				key := rand.Intn(1000) // Generate random key
				var success bool

				// Determine which list type we're using and perform the operation
				switch l := list.(type) {
				case *LinkedList:
					if insert {
						success = l.Insert(key) // Insert into global lock list
					} else {
						success = l.Lookup(key) // Lookup in global lock list
					}
				case *HLinkedList:
					if insert {
						success = l.Insert(key) // Insert into hand-over-hand list
					} else {
						success = l.Lookup(key) // Lookup in hand-over-hand list
					}
				}

				if success {
					localSuccess++
				}
			}

			// Update global success count safely
			mutex := &sync.Mutex{}
			mutex.Lock()
			successCount += localSuccess
			mutex.Unlock()
		}()
	}

	wg.Wait() // Wait for all threads to complete
	elapsed := time.Since(start) // Calculate elapsed time
	fmt.Printf("Time elapsed: %v | Successful operations: %d\n", elapsed, successCount)
}

// main function runs benchmarks for both linked list implementations
func main() {
	fmt.Println("Benchmarking Global Lock List")
	globalList := &LinkedList{}
	benchmark(globalList, true)  // Measure insert performance
	benchmark(globalList, false) // Measure lookup performance

	fmt.Println("Benchmarking Hand-over-Hand Lock List")
	handList := &HLinkedList{}
	benchmark(handList, true)  // Measure insert performance
	benchmark(handList, false) // Measure lookup performance
}
