package main

import "sync"

// Node represents an element in the linked list
type Node struct {
	key  int
	next *Node
}

// LinkedList uses a single global lock for synchronization
type LinkedList struct {
	head *Node      // Pointer to the first node
	lock sync.Mutex // Mutex for thread-safe operations
}

// Insert adds a new node to the beginning of the list
// Returns true if insertion is successful, false if the key already exists
func (l *LinkedList) Insert(key int) bool {
	l.lock.Lock()          // Lock the entire list
	defer l.lock.Unlock()  // Ensure unlock happens at function exit

	// Check if key already exists in the list
	curr := l.head
	for curr != nil {
		if curr.key == key {
			return false // Key already exists, insertion fails
		}
		curr = curr.next
	}

	// Insert new node at the head
	newNode := &Node{key: key, next: l.head}
	l.head = newNode
	return true // Successfully inserted
}

// Lookup searches for a key in the list
// Returns true if the key is found, false otherwise
func (l *LinkedList) Lookup(key int) bool {
	l.lock.Lock()         // Lock the entire list
	defer l.lock.Unlock() // Ensure unlock happens at function exit

	curr := l.head
	for curr != nil {
		if curr.key == key {
			return true // Key found
		}
		curr = curr.next
	}
	return false // Key not found
}
