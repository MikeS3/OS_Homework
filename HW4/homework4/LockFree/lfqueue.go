package main

import (
	"sync"
)

type Node struct {
	value interface{}
	next  *Node
}

type Queue struct {
	Head *Node
	Tail *Node
	mu   sync.Mutex
}

type NodeT struct {
	ptr   *Node
	count uint64
}

// Initialize creates a new queue with a single dummy node
func (Q *Queue) initialize() {
	node := &Node{} // Allocate a new node
	node.next = nil // The next pointer is null
	Q.Head = node   // Both Head and Tail point to this node
	Q.Tail = node
}

// Enqueue adds a new value to the end of the queue
func (Q *Queue) enqueue(value interface{}) {
	for {
		Q.mu.Lock()
		tail := Q.Tail
		next := tail.next

		if tail == Q.Tail {
			if next == nil {
				node := &Node{value: value}
				tail.next = node
				Q.Tail = node
				Q.mu.Unlock()
				return
			}
			Q.Tail = next
		}
		Q.mu.Unlock()
	}
}

// Dequeue removes and returns the value from the front of the queue
func (Q *Queue) dequeue() (interface{}, bool) {
	for {
		Q.mu.Lock()
		head := Q.Head
		tail := Q.Tail
		next := head.next

		if head == Q.Head {
			if head == tail {
				if next == nil {
					Q.mu.Unlock()
					return nil, false // Queue is empty
				}
				Q.Tail = next
			}

			value := next.value
			Q.Head = next
			Q.mu.Unlock()

			// Optionally free the node
			return value, true
		}
		Q.mu.Unlock()
	}
}
