package main

import "sync"

// HNode represents a node in the hand-over-hand locking linked list
type HNode struct {
	key  int
	next *HNode
	lock sync.Mutex // Each node has its own lock
}

// HLinkedList represents a linked list using hand-over-hand locking
type HLinkedList struct {
	head     *HNode    // Pointer to the first node
	listLock sync.Mutex // Lock for head modification
}

// Insert adds a new node to the beginning of the list
// Returns true if insertion is successful, false if the key already exists
func (l *HLinkedList) Insert(key int) bool {
	l.listLock.Lock()      // Lock list to safely modify the head
	defer l.listLock.Unlock()

	// Check if key already exists before inserting
	curr := l.head
	for curr != nil {
		curr.lock.Lock() // Lock each node while traversing
		if curr.key == key {
			curr.lock.Unlock()
			return false // Key already exists
		}
		curr.lock.Unlock()
		curr = curr.next
	}

	// Insert new node at the head
	newNode := &HNode{key: key, next: l.head}
	l.head = newNode
	return true // Successfully inserted
}

// Lookup searches for a key in the list using hand-over-hand locking
// Returns true if the key is found, false otherwise
func (l *HLinkedList) Lookup(key int) bool {
	l.listLock.Lock() // Lock the list before accessing head
	curr := l.head
	if curr != nil {
		curr.lock.Lock() // Lock the first node
	}
	l.listLock.Unlock()

	// Traverse the list using hand-over-hand locking
	for curr != nil {
		if curr.key == key {
			curr.lock.Unlock()
			return true // Key found
		}

		next := curr.next
		if next != nil {
			next.lock.Lock() // Lock the next node before unlocking the current one
		}
		curr.lock.Unlock() // Unlock the current node
		curr = next
	}

	return false // Key not found
}
