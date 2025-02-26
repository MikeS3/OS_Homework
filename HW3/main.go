package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numOperations = 1000
	numThreads    = 10
)

// Benchmark function
func benchmark(list interface{}, insert bool) {
	var wg sync.WaitGroup
	start := time.Now()

	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < numOperations; j++ {
				key := rand.Intn(1000)
				switch l := list.(type) {
				case *LinkedList:
					if insert {
						l.Insert(key)
					} else {
						l.Lookup(key)
					}
				case *HLinkedList:
					if insert {
						l.Insert(key)
					} else {
						l.Lookup(key)
					}
				}
			}
		}()
	}

	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Time elapsed: %v\n", elapsed)
}

func main() {
	fmt.Println("Benchmarking Global Lock List")
	globalList := &LinkedList{}
	benchmark(globalList, true)
	benchmark(globalList, false)

	fmt.Println("Benchmarking Hand-over-Hand Lock List")
	handList := &HLinkedList{}
	benchmark(handList, true)
	benchmark(handList, false)
}
