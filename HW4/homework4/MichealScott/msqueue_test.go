package main

import (
	"testing"
)

// BenchmarkQueue measures the performance of enqueue and dequeue operations
func BenchmarkQueue(b *testing.B) {
	q := NewQueue()

	b.ResetTimer() // Reset the timer to exclude setup time

	for i := 0; i < b.N; i++ { // Go automatically determines b.N
		q.Enqueue(i)
		q.Dequeue()
	}
}
