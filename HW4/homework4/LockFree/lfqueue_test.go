package main

import (
	"testing"
)

func BenchmarkQueue_Enqueue(b *testing.B) {
	// Create a new queue
	var q Queue
	q.initialize()

	// Run the benchmark for the number of iterations
	for i := 0; i < b.N; i++ {
		q.enqueue(i)
	}
}

func BenchmarkQueue_Dequeue(b *testing.B) {
	// Create a new queue and enqueue some values
	var q Queue
	q.initialize()
	for i := 0; i < 1000; i++ {
		q.enqueue(i)
	}

	// Run the benchmark for the number of iterations
	b.ResetTimer() // Reset timer to ignore setup time (initial enqueues)
	for i := 0; i < b.N; i++ {
		q.dequeue()
	}
}
