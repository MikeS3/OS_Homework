package main

import "sync"

// Node structure with its own lock
type HNode struct {
	key  int
	next *HNode
	lock sync.Mutex
}

// Hand-over-Hand Linked List
type HLinkedList struct {
	head     *HNode
	listLock sync.Mutex
}

// Insert a new node at the beginning
func (l *HLinkedList) Insert(key int) {
	l.listLock.Lock()
	defer l.listLock.Unlock()

	newNode := &HNode{key: key, next: l.head}
	l.head = newNode
}

// Lookup a key using hand-over-hand locking
func (l *HLinkedList) Lookup(key int) bool {
	l.listLock.Lock()
	curr := l.head
	if curr != nil {
		curr.lock.Lock()
	}
	l.listLock.Unlock()

	for curr != nil {
		if curr.key == key {
			curr.lock.Unlock()
			return true
		}

		next := curr.next
		if next != nil {
			next.lock.Lock()
		}
		curr.lock.Unlock()
		curr = next
	}

	return false
}
