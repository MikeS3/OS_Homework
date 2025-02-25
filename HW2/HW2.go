//Code for Ticket lock/CAS Spin lock and benchmark generated by ChatGPT
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type TicketLock struct {
	ticket int32
	turn   int32
}

func (l *TicketLock) Lock() {
	myTurn := atomic.AddInt32(&l.ticket, 1) - 1 // FetchAndAdd equivalent
	for atomic.LoadInt32(&l.turn) != myTurn {
		// Spin-wait
	}
}

func (l *TicketLock) Unlock() {
	atomic.AddInt32(&l.turn, 1)
}

type CASSpinLock struct {
	flag int32
}

func (l *CASSpinLock) Lock() {
	for !atomic.CompareAndSwapInt32(&l.flag, 0, 1) {
		// Spin-wait
	}
}

func (l *CASSpinLock) Unlock() {
	atomic.StoreInt32(&l.flag, 0)
}

// Function to benchmark Ticket Lock
func benchmarkTicketLock(numGoroutines int) time.Duration {
	var lock TicketLock
	var wg sync.WaitGroup
	start := time.Now()

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			lock.Lock()
			time.Sleep(10 * time.Millisecond) // Simulate some work
			lock.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()

	return time.Since(start)
}

// Function to benchmark CAS Spin Lock
func benchmarkCASSpinLock(numGoroutines int) time.Duration {
	var lock CASSpinLock
	var wg sync.WaitGroup
	start := time.Now()

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			lock.Lock()
			time.Sleep(10 * time.Millisecond) // Simulate some work
			lock.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()

	return time.Since(start)
}

func main() {
	numGoroutines := 32 // Adjust this for different contention levels

	// Benchmark Ticket Lock
	ticketLockTime := benchmarkTicketLock(numGoroutines)
	fmt.Printf("Ticket Lock (with %d goroutines) took: %v\n", numGoroutines, ticketLockTime)

	// Benchmark CAS Spin Lock
	casLockTime := benchmarkCASSpinLock(numGoroutines)
	fmt.Printf("CAS Spin Lock (with %d goroutines) took: %v\n", numGoroutines, casLockTime)
}
