package main

import (
	"sync"
)

type Node struct {
	value int
	next  *Node
}

type Queue struct {
	head     *Node
	tail     *Node
	headLock sync.Mutex
	tailLock sync.Mutex
}

func NewQueue() *Queue {
	dummy := &Node{}
	return &Queue{
		head: dummy,
		tail: dummy,
	}
}

func (q *Queue) Enqueue(value int) {
	tmp := &Node{value: value}

	q.tailLock.Lock()
	q.tail.next = tmp
	q.tail = tmp
	q.tailLock.Unlock()
}

func (q *Queue) Dequeue() (int, bool) {
	q.headLock.Lock()
	tmp := q.head
	newHead := tmp.next
	if newHead == nil {
		q.headLock.Unlock()
		return 0, false // Queue was empty
	}

	value := newHead.value
	q.head = newHead
	q.headLock.Unlock()
	return value, true
}
