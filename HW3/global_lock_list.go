package main

import "sync"

// Node structure
type Node struct {
	key  int
	next *Node
}

// LinkedList with a global lock
type LinkedList struct {
	head *Node
	lock sync.Mutex
}

// Insert a new node at the beginning
func (l *LinkedList) Insert(key int) {
	l.lock.Lock()
	defer l.lock.Unlock()

	newNode := &Node{key: key, next: l.head}
	l.head = newNode
}

// Lookup a key in the list
func (l *LinkedList) Lookup(key int) bool {
	l.lock.Lock()
	defer l.lock.Unlock()

	curr := l.head
	for curr != nil {
		if curr.key == key {
			return true
		}
		curr = curr.next
	}
	return false
}
